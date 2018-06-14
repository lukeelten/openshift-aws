package main

import "../../internal/orchestration"


const default_output = "/app/generated"
const default_base = "/app/"

func main() {
	oc := orchestration.NewOrchestration(default_output, default_base)

	// Check Directories
	oc.Validate()

	oc.HandleFlags()
	oc.RunTerraform()
	oc.GenerateConfiguration()
	oc.RunInstaller()
	oc.RunPostInstallationConfig()
}
