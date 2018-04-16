package main

import (
	"settings"
	"os"
	"aws"
	"openshift"
	"ansible"
)

const GEN_DIR = "generated/"
const INVENTORY = GEN_DIR + "inventory"
const SSH_CONFIG_FILE = GEN_DIR + "ssh.cfg"


type Test struct {
	EfsId string
	Region string
}

func main() {

	settings.ParseFlags()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	awsConfig := aws.NewConfig("eu-central-1", "", "")
	awsConfig.InitSession()

	sshConfig := openshift.GenerateSshConfig()
	sshConfig.WriteConfig(SSH_CONFIG_FILE)

	config := openshift.GenerateConfig(SSH_CONFIG_FILE)
	config.GenerateInventory(INVENTORY)

	installerPath := wd + "/../openshift-ansible"

	ansible.CheckReadiness(INVENTORY)

	playbook := ansible.OpenPlaybook(installerPath + "/playbooks/byo/config.yml")
	playbook.Run(INVENTORY)

	/*
	playbook = ansible.OpenPlaybook(wd + "/nfs-setup.yml")
	playbook.Run(INVENTORY)
	*/
}