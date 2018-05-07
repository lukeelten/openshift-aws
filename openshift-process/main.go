package main

import (
	"configuration"
	"os"
	"openshift"
	"util"
	"terraform"
	"fmt"
	"ansible"
	"aws"
	"time"
	"sync"
)

const GEN_DIR = "generated/"
const INVENTORY = GEN_DIR + "inventory"
const SSH_CONFIG_FILE = GEN_DIR + "ssh.cfg"
const SSH_KEY_FILE = "ssh.key"

func main() {
	var config *configuration.InputVars

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

	if !cmdFlags.SkipTerraform {
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
	}

	aws.InitSession(config)

	var wg sync.WaitGroup
	if !cmdFlags.SkipConfig {
		wg.Add(3)
		go generateSshConfig(config, &wg)
		go generatePersistenceConfig(config, &wg)
		go generateInventory(config, &wg)
	}

	if !cmdFlags.SkipTerraform {
		// three minutes should be enough to init EC2 instances
		fmt.Println("\nWaiting for instances to get ready ...")
		time.Sleep(3 * time.Minute)
	}

	wg.Wait() // wait for go routines to finish, should be done by now, but to be sure ...

	installerPath := wd + "/../openshift-ansible"

	/*
	playbook := ansible.OpenPlaybook(installerPath + "/playbooks/byo/config.yml")
	if err := playbook.Run(INVENTORY); err != nil {
		util.ExitOnError("Failed to run OpenShift installer.", err)
	}
	*/

	playbook := ansible.OpenPlaybook(installerPath + "/playbooks/prerequisites.yml")
	if err := playbook.Run(INVENTORY); err != nil {
		util.ExitOnError("Failed to run OpenShift prerequisites.", err)
	}

	playbook = ansible.OpenPlaybook(installerPath + "/playbooks/deploy_cluster.yml")
	if err := playbook.Run(INVENTORY); err != nil {
		util.ExitOnError("Failed to run OpenShift installer.", err)
	}

	/*
	fmt.Println("\nWaiting for OpenShift to become ready ...")
	time.Sleep(2 * time.Minute)

	playbook = ansible.OpenPlaybook(wd + "/playbooks/post-config.yml")
	if err := playbook.Run(INVENTORY); err != nil {
		util.ExitOnError("Failed to run post installation configuration", err)
	}
	*/
}

func generateSshConfig(config *configuration.InputVars, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	sshConfig := openshift.GenerateSshConfig(config)
	if err := sshConfig.WriteConfig(SSH_CONFIG_FILE); err != nil {
		util.ExitOnError("Cannot write SSH configuration file", err)
	}
}

func generatePersistenceConfig(config *configuration.InputVars, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	persistenceConfig := openshift.NewPersistenceConfig(config)
	if err := persistenceConfig.GeneratePersistenceConfigFiles(GEN_DIR); err != nil {
		util.ExitOnError("Cannot write persistence storage configuration.", err)
	}
}

func generateInventory(config *configuration.InputVars, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	openshiftConfig := openshift.GenerateConfig(SSH_CONFIG_FILE, config)
	if err:= openshiftConfig.GenerateInventory(INVENTORY); err != nil {
		util.ExitOnError("Cannot write OpenShift inventory file.", err)
	}
}