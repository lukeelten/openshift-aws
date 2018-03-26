package main

import (
	"settings"
	"os"
	"ansible"
	"aws"
	"terraform"
)

const INVENTORY = "myinventory"

func main() {
	settings.ParseFlags()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	fullPath := wd + "/" + settings.ActiveSettings.TerraformDir

	terraform.InitTerraform(fullPath)
	// TODO: terraform operations

	aws.InitAws()

	inventory := ansible.GenerateOpenshiftInventory(INVENTORY)
	inventory.Write()

	/*
	playbook := ansible.OpenPlaybook("/home/lukeelten/Projekte/codecentric/repo/openshift-ansible/playbooks/prerequisites.yml")
	playbook.Run(INVENTORY)

	playbook = ansible.OpenPlaybook("/home/lukeelten/Projekte/codecentric/repo/openshift-ansible/playbooks/deploy_cluster.yml")
	playbook.Run(INVENTORY)
	*/

	playbook := ansible.OpenPlaybook("/home/lukeelten/Projekte/codecentric/repo/openshift-ansible/playbooks/byo/config.yml")
	playbook.Run(INVENTORY)

	ansible.ExecuteRemote(INVENTORY, "masters", "/bin/oadm policy add-cluster-role-to-user cluster-admin admin")
}
