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
)

const GEN_DIR = "generated/"
const INVENTORY = GEN_DIR + "inventory"
const SSH_CONFIG_FILE = GEN_DIR + "ssh.cfg"

func main() {

	settings := configuration.ParseFlags()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	terraformDir := wd + "/../terraform"
	tf := terraform.NewConfig(terraformDir, &settings)
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

	// three minutes should be enough
	time.Sleep(3 * time.Minute)

	settings.AWSConfig.InitSession()

	sshConfig := openshift.GenerateSshConfig()
	if err := sshConfig.WriteConfig(SSH_CONFIG_FILE); err != nil {
		util.ExitOnError("Cannot write SSH configuration file", err)
	}

	config := openshift.GenerateConfig(SSH_CONFIG_FILE)
	if err:= config.GenerateInventory(INVENTORY); err != nil {
		util.ExitOnError("Cannot write OpenShift inventory file.", err)
	}

	persistenceConfig := openshift.NewPersistenceConfig(&settings)
	if err := persistenceConfig.GeneratePersistenceConfigFiles(GEN_DIR); err != nil {
		util.ExitOnError("Cannot write persistence storage configuration.", err)
	}

	installerPath := wd + "/../openshift-ansible"

	//ansible.CheckReadiness(INVENTORY)

	playbook := ansible.OpenPlaybook(installerPath + "/playbooks/byo/config.yml")
	if err := playbook.Run(INVENTORY); err != nil {
		util.ExitOnError("Failed to run OpenShift installer.", err)
	}

	playbook = ansible.OpenPlaybook(wd + "/playbooks/post-config.yml")
	if err := playbook.Run(INVENTORY); err != nil {
		util.ExitOnError("Failed to run post installation configuration", err)
	}
}