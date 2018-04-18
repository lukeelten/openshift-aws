package main

import (
	"configuration"
	"os"
	"openshift"
	"ansible"
	"terraform"
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
	tf.InitTerraform()
	tf.Apply()

	// @todo wait for AWS to become ready

	settings.AWSConfig.InitSession()

	sshConfig := openshift.GenerateSshConfig()
	sshConfig.WriteConfig(SSH_CONFIG_FILE)

	config := openshift.GenerateConfig(SSH_CONFIG_FILE)
	config.GenerateInventory(INVENTORY)

	persistenceConfig := openshift.NewPersistenceConfig(&settings)
	persistenceConfig.GeneratePersistenceConfigFiles(GEN_DIR)

	installerPath := wd + "/../openshift-ansible"

	ansible.CheckReadiness(INVENTORY)

	playbook := ansible.OpenPlaybook(installerPath + "/playbooks/byo/config.yml")
	playbook.Run(INVENTORY)

	playbook = ansible.OpenPlaybook(wd + "/playbooks/post-config.yml")
	playbook.Run(INVENTORY)
}