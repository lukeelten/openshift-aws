package main

import (
	"configuration"
	"os"
	"openshift"
	"util"
	"terraform"
	"time"
	"fmt"
	"ansible"
	"aws"
)

const GEN_DIR = "generated/"
const INVENTORY = GEN_DIR + "inventory"
const SSH_CONFIG_FILE = GEN_DIR + "ssh.cfg"
const SSH_KEY_FILE = "ssh.key"

var config *configuration.InputVars

func main() {

	cmdFlags := configuration.ParseFlags()
	if len(cmdFlags.ConfigFile) > 0 {
		config = configuration.LoadInputVars(cmdFlags.ConfigFile)
	} else {
		config = configuration.DefaultConfig()
	}
	config.MergeCmdFlags(cmdFlags)
	config.Validate()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	key := util.NewKeyPair()
	key.WritePrivateKey(GEN_DIR + SSH_KEY_FILE)
	key.WritePublicPem(GEN_DIR + SSH_KEY_FILE + ".pub")
	agent := util.NewSshAgentClient()
	agent.AddKey(key)

	terraformDir := wd + "/../terraform"
	tf := terraform.NewConfig(terraformDir, key.GetPublicKey(), config)
	if !tf.InitTerraform() {
		panic("Cannot init terraform. Is the directory correct? " + terraformDir)
	}

	if err := tf.Validate(); err != nil {
		util.ExitOnError("Invalid terraform configuration.", err)
	}

	if err := tf.Apply(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during terraform process. Remaining infrastructure will be destroyed.")
		tf.Destroy()
		util.ExitOnError("Error during terraform process", err)
	}

	aws.InitSession(config)

	go generateSshConfig()
	go generatePersistenceConfig()
	go generateInventory()


	// three minutes should be enough to init EC2 instances
	fmt.Println("\nWaiting for instances to get ready ...")
	time.Sleep(3 * time.Minute)

	installerPath := wd + "/../openshift-ansible"

	playbook := ansible.OpenPlaybook(installerPath + "/playbooks/byo/config.yml")
	if err := playbook.Run(INVENTORY); err != nil {
		util.ExitOnError("Failed to run OpenShift installer.", err)
	}

	playbook = ansible.OpenPlaybook(wd + "/playbooks/post-config.yml")
	if err := playbook.Run(INVENTORY); err != nil {
		util.ExitOnError("Failed to run post installation configuration", err)
	}
}

func generateSshConfig() {
	sshConfig := openshift.GenerateSshConfig()
	if err := sshConfig.WriteConfig(SSH_CONFIG_FILE); err != nil {
		util.ExitOnError("Cannot write SSH configuration file", err)
	}
}

func generatePersistenceConfig() {
	persistenceConfig := openshift.NewPersistenceConfig(config)
	if err := persistenceConfig.GeneratePersistenceConfigFiles(GEN_DIR); err != nil {
		util.ExitOnError("Cannot write persistence storage configuration.", err)
	}
}

func generateInventory() {
	config := openshift.GenerateConfig(SSH_CONFIG_FILE)
	if err:= config.GenerateInventory(INVENTORY); err != nil {
		util.ExitOnError("Cannot write OpenShift inventory file.", err)
	}
}