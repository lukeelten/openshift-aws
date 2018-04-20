package main

import (
	"configuration"
	"os"
	"openshift"
	"ansible"
	"util"
	"terraform"
	"time"
	"fmt"
)

const GEN_DIR = "generated/"
const INVENTORY = GEN_DIR + "inventory"
const SSH_CONFIG_FILE = GEN_DIR + "ssh.cfg"

type Runner func ()

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

	// @todo find better solution
	time.Sleep(5 * time.Minute)

	trackTime("Init AWS Session", func() {
		settings.AWSConfig.InitSession()
	})

	trackTime("Generate SSH Config", func() {
		sshConfig := openshift.GenerateSshConfig()
		if err := sshConfig.WriteConfig(SSH_CONFIG_FILE); err != nil {
			util.ExitOnError("Cannot write SSH configuration file", err)
		}
	})

	trackTime("Generate OpenShift Inventory", func() {
		config := openshift.GenerateConfig(SSH_CONFIG_FILE)
		if err:= config.GenerateInventory(INVENTORY); err != nil {
			util.ExitOnError("Cannot write OpenShift inventory file.", err)
		}
	})

	trackTime("Generate Persistence Configuration", func() {
		persistenceConfig := openshift.NewPersistenceConfig(&settings)
		if err := persistenceConfig.GeneratePersistenceConfigFiles(GEN_DIR); err != nil {
			util.ExitOnError("Cannot write persistence storage configuration.", err)
		}
	})

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

func printTime(msg string, started time.Time) {
	elapsed := time.Since(started)
	fmt.Printf("%s: %s", msg, elapsed)
}

func trackTime(msg string, function Runner) {
	defer printTime(msg, time.Now())
	function()
}