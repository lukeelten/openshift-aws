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


type OrchestrationConfig struct {
	GenDir string
	Inventory string
	SshKeyFile string
	SshConfigFile string

	wd string

	cmdFlags configuration.CmdFlags
	config *configuration.InputVars
}

func NewOrchestration(genDir, inventory, sshKeyFile, sshConfigFile string) OrchestrationConfig {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return OrchestrationConfig{
		GenDir: genDir,
		Inventory: inventory,
		SshKeyFile: sshKeyFile,
		SshConfigFile: sshConfigFile,
		wd: wd,
	}
}


func (oc OrchestrationConfig) HandleFlags() {
	oc.cmdFlags = configuration.ParseFlags()
	if len(oc.cmdFlags.ConfigFile) > 0 {
		oc.config = configuration.LoadInputVars(oc.cmdFlags.ConfigFile)
	} else {
		oc.config = configuration.DefaultConfig()
	}

	oc.config.MergeCmdFlags(oc.cmdFlags)
	util.ExitOnError("Invalid configuration found", oc.config.Validate())
}


func (oc OrchestrationConfig) RunTerraform() {
	if !oc.cmdFlags.SkipTerraform {
		key := util.NewKeyPair()
		key.WritePrivateKey(oc.GenDir + oc.SshKeyFile)
		key.WritePublicPem(oc.GenDir + oc.SshKeyFile + ".pub")
		agent := util.NewSshAgentClient()
		agent.AddKey(key)

		terraformDir := oc.wd + "/../terraform"
		tf := terraform.NewConfig(terraformDir, key.GetPublicKey(), oc.config)
		tf.GenerateVarsFile()
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

func (oc OrchestrationConfig) GenerateConfiguration() {
	aws.InitSession(oc.config)

	var wg sync.WaitGroup
	if !oc.cmdFlags.SkipConfig {
		wg.Add(3)
		go generateSshConfig(&oc, &wg)
		go generatePersistenceConfig(&oc, &wg)
		go generateInventory(&oc, &wg)
	}

	if !oc.cmdFlags.SkipTerraform {
		// three minutes should be enough to init EC2 instances
		fmt.Println("\nWaiting for instances to get ready ...")
		time.Sleep(3 * time.Minute)
	}

	wg.Wait() // wait for go routines to finish, should be done by now, but to be sure ...
}

func (oc OrchestrationConfig) RunInstaller() {
	installerPath := oc.wd + "/openshift-ansible"

	var playbook *ansible.Playbook
	if !oc.cmdFlags.SkipPre {
		playbook = ansible.OpenPlaybook(installerPath + "/playbooks/prerequisites.yml")
		err := playbook.Run(oc.Inventory)
		util.ExitOnError("Failed to run OpenShift prerequisites.", err)
	}

	playbook = ansible.OpenPlaybook(installerPath + "/playbooks/deploy_cluster.yml")
	err := playbook.Run(oc.Inventory)
	util.ExitOnError("Failed to run OpenShift installer.", err)

	fmt.Println("\nWaiting for OpenShift to become ready ...")
	time.Sleep(2 * time.Minute)
}

func (oc OrchestrationConfig) RunPostInstallationConfig() {
	playbook := ansible.OpenPlaybook(oc.wd + "/playbooks/post-config.yml")
	err := playbook.Run(oc.Inventory)
	util.ExitOnError("Failed to run post installation configuration", err)

	if oc.config.Storage.EnableEfs {
		playbook = ansible.OpenPlaybook(oc.wd + "/playbooks/efs.yml")
		err := playbook.Run(oc.Inventory)
		util.ExitOnError("Failed to run EFS configuration", err)
	}

	if oc.config.Storage.EnableEbs {
		playbook = ansible.OpenPlaybook(oc.wd + "/playbooks/ebs.yml")
		err := playbook.Run(oc.Inventory)
		util.ExitOnError("Failed to run EBS configuration", err)
	}
}

func generateSshConfig(oc *OrchestrationConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	sshConfig := openshift.GenerateSshConfig(oc.config)
	err := sshConfig.WriteConfig(oc.SshConfigFile)
	util.ExitOnError("Cannot write SSH configuration file", err)
}

func generatePersistenceConfig(oc *OrchestrationConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	persistenceConfig := openshift.NewPersistenceConfig(oc.config)
	err := persistenceConfig.GeneratePersistenceConfigFiles(oc.GenDir)
	util.ExitOnError("Cannot write persistence storage configuration.", err)

}

func generateInventory(oc *OrchestrationConfig, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	openshiftConfig := openshift.GenerateConfig(oc.SshConfigFile, oc.config)
	err:= openshiftConfig.GenerateInventory(oc.Inventory)
	util.ExitOnError("Cannot write OpenShift inventory file.", err)
}