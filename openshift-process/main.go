package main

import (
	"settings"
	"os"
	"ansible"
	"aws"
)

const INVENTORY = "myinventory"

func main() {
	settings.ParseFlags()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	aws.InitAws()

	inventory := ansible.GenerateOpenshiftInventory(INVENTORY)
	inventory.Write()

	installerPath := wd + "/../openshift-ansible"

	ansible.CheckReadiness(INVENTORY)

	playbook := ansible.OpenPlaybook(installerPath + "/playbooks/byo/config.yml")
	playbook.Run(INVENTORY)

	/*
	playbook = ansible.OpenPlaybook(wd + "/nfs-setup.yml")
	playbook.Run(INVENTORY)
	*/
}

/*
	playbook := ansible.OpenPlaybook("/home/lukeelten/Projekte/codecentric/repo/openshift-ansible/playbooks/byo/config.yml")
	playbook.Run(INVENTORY)
 */