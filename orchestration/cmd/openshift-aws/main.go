package main

import "../../internal/orchestration"

const GEN_DIR = "generated/"
const INVENTORY = GEN_DIR + "inventory"
const SSH_CONFIG_FILE = GEN_DIR + "ssh.cfg"
const SSH_KEY_FILE = "ssh.key"
const TERRAFORM_STATE = "terraform.tfstate"

func main() {

	oc := orchestration.NewOrchestration(
		GEN_DIR,
		INVENTORY,
		SSH_KEY_FILE,
		SSH_CONFIG_FILE,
		GEN_DIR + TERRAFORM_STATE)

	oc.HandleFlags()
	oc.RunTerraform()
	oc.GenerateConfiguration()
	oc.RunInstaller()
	oc.RunPostInstallationConfig()
}
