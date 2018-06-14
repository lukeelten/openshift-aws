package orchestration

import (
	"os"
	"fmt"
	"sync"
	"time"
	"../configuration"
	"../util"
	"../terraform"
	"../aws"
	"../ansible"
	"../openshift"
)

// Time to wait until EC2 instances are ready (in minutes)
const WAIT_TIME = 4

type OrchestrationConfig struct {
	OutputDir string
	BaseDir string

	terraformDir string
	installerDir string
	playbooksDir string
	templatesDir string

	Inventory string
	SshKeyFile string
	SshConfigFile string
	TerraformState string
	TerraformConfig string

	cmdFlags configuration.CmdFlags
	config *configuration.InputVars
}

func NewOrchestration(outputDir, baseDir string) *OrchestrationConfig {
	outputDir = prettyPrint(outputDir)
	baseDir = prettyPrint(baseDir)

	return &OrchestrationConfig{
		OutputDir: prettyPrint(outputDir),
		BaseDir: baseDir,

		terraformDir: baseDir + "terraform/",
		installerDir: baseDir + "openshift-ansible/",
		playbooksDir: baseDir + "playbooks/",
		templatesDir: baseDir + "templates/",

		Inventory: default_inventory,
		SshKeyFile: default_ssh_key,
		SshConfigFile: default_ssh_config,
		TerraformState: default_terraform_state,
		TerraformConfig: default_terraform_config,
	}
}

func (oc *OrchestrationConfig) Validate() {
	if !checkTerraformDir(oc.terraformDir) {
		panic("Invalid Terraform directory detected. Check the path: " + oc.terraformDir)
	}

	if !checkInstallerDir(oc.installerDir) {
		panic("Invalid Installer directory detected. Check the path: " + oc.installerDir)
	}

	if  !checkOutputDir(oc.OutputDir) {
		panic("Invalid Output directory detected. Check the path: " + oc.OutputDir)
	}

	if !checkPlaybookDir(oc.playbooksDir) {
		panic("Invalid Playbooks directory detected. Check the path: " + oc.playbooksDir)
	}

	if !checkTemplatesDir(oc.templatesDir) {
		panic("Invalid Templates directory detected. Check the path: " + oc.templatesDir)
	}
}

func (oc *OrchestrationConfig) HandleFlags() {
	oc.cmdFlags = configuration.ParseFlags()
	if len(oc.cmdFlags.ConfigFile) > 0 {
		oc.config = configuration.LoadInputVars(oc.cmdFlags.ConfigFile)
	} else {
		oc.config = configuration.DefaultConfig()
	}

	oc.config.MergeCmdFlags(oc.cmdFlags)
	util.ExitOnError("Invalid configuration found", oc.config.Validate())
}


func (oc *OrchestrationConfig) RunTerraform() {
	if !oc.cmdFlags.SkipTerraform {
		key := util.NewKeyPair()
		key.WritePrivateKey(oc.OutputDir + oc.SshKeyFile)
		key.WritePublicPem(oc.OutputDir + oc.SshKeyFile + ".pub")
		agent := util.NewSshAgentClient()
		agent.AddKey(key)

		terraformDir := oc.terraformDir
		tf := terraform.NewConfig(terraformDir, oc.OutputDir + oc.TerraformState, key.GetPublicKey(), oc.config)
		tf.GenerateVarsFile(oc.OutputDir + oc.TerraformConfig)
		err := tf.InitTerraform()
		util.ExitOnError("Cannot init terraform. Is the directory correct? " + terraformDir, err)

		err = tf.Validate()
		util.ExitOnError("Invalid terraform configuration.", err)

		if err := tf.Apply(); err != nil {
			fmt.Fprintf(os.Stderr, "Error during terraform process. Remaining infrastructure will be destroyed.")
			tf.Destroy()
			util.ExitOnError("Error during terraform process", err)
		}
	}
}

func (oc *OrchestrationConfig) GenerateConfiguration() {
	aws.InitSession(oc.config)

	var wg sync.WaitGroup
	if !oc.cmdFlags.SkipConfig {
		wg.Add(3)
		go generateSshConfig(oc, &wg)
		go generatePersistenceConfig(oc, &wg)
		go generateInventory(oc, &wg)
	}

	if !oc.cmdFlags.SkipTerraform {
		// three minutes should be enough to init EC2 instances
		fmt.Println("\nWaiting for instances to get ready ...")
		time.Sleep(WAIT_TIME * time.Minute)
	}

	wg.Wait() // wait for go routines to finish, should be done by now, but to be sure ...
}

func (oc *OrchestrationConfig) RunInstaller() {
	installerPath := oc.installerDir

	var playbook *ansible.Playbook
	if !oc.cmdFlags.SkipPre {
		playbook = ansible.OpenPlaybook(installerPath + "playbooks/prerequisites.yml")
		err := playbook.Run(oc.OutputDir + oc.Inventory)
		util.ExitOnError("Failed to run OpenShift prerequisites.", err)
	}

	playbook = ansible.OpenPlaybook(installerPath + "playbooks/deploy_cluster.yml")
	err := playbook.Run(oc.OutputDir + oc.Inventory)
	util.ExitOnError("Failed to run OpenShift installer.", err)

	fmt.Println("\nWaiting for OpenShift to become ready ...")
	time.Sleep(2 * time.Minute)
}

func (oc *OrchestrationConfig) RunPostInstallationConfig() {
	playbook := ansible.OpenPlaybook(oc.playbooksDir + "post-config.yml")
	err := playbook.Run(oc.OutputDir + oc.Inventory)
	util.ExitOnError("Failed to run post installation configuration", err)

	if oc.config.Storage.EnableEfs {
		playbook = ansible.OpenPlaybook(oc.playbooksDir + "/efs.yml")
		err := playbook.Run(oc.OutputDir + oc.Inventory)
		util.ExitOnError("Failed to run EFS configuration", err)
	}

	if oc.config.Storage.EnableEbs {
		playbook = ansible.OpenPlaybook(oc.playbooksDir + "/ebs.yml")
		err := playbook.Run(oc.OutputDir + oc.Inventory)
		util.ExitOnError("Failed to run EBS configuration", err)
	}
}

func generateSshConfig(oc *OrchestrationConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	sshConfig := openshift.GenerateSshConfig(oc.config)
	err := sshConfig.WriteConfig(oc.OutputDir + oc.SshConfigFile)
	util.ExitOnError("Cannot write SSH configuration file", err)
}

func generatePersistenceConfig(oc *OrchestrationConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	persistenceConfig := openshift.NewPersistenceConfig(oc.config)
	err := persistenceConfig.GeneratePersistenceConfigFiles(oc.OutputDir)
	util.ExitOnError("Cannot write persistence storage configuration.", err)

}

func generateInventory(oc *OrchestrationConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	openshiftConfig := openshift.GenerateConfig(oc.OutputDir + oc.SshConfigFile, oc.config)
	err:= openshiftConfig.GenerateInventory(oc.OutputDir + oc.Inventory)
	util.ExitOnError("Cannot write OpenShift inventory file.", err)
}