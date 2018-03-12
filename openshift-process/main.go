package main

import (
	"aws"
	"ansible"
)

const INVENTORY = "myinventory"

func main() {
	aws.InitAws()

	inventory := ansible.GenerateOpenshiftInventory(INVENTORY)
	inventory.Write()

	playbook := ansible.OpenPlaybook("/home/lukeelten/Projekte/codecentric/openshift-ansible/playbooks/byo/config.yml")
	playbook.Run(INVENTORY)

	ansible.ExecuteRemote(INVENTORY, "masters", "/bin/oadm policy add-cluster-role-to-user cluster-admin admin")
}
