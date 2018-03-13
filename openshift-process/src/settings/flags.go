package settings

import "flag"

type Settings struct {
	NumMasters uint
	NumInfra uint
	NumApplications uint
	TerraformDir string
}

var ActiveSettings Settings

type flags struct {
	masters *uint
	infra *uint
	app *uint
	containerized *bool
	terraform *string
}

var cmdFlags flags

func initFlags() {
	cmdFlags.masters = flag.Uint("master-nodes", 2, "Number of master nodes")
	cmdFlags.infra = flag.Uint("infra-nodes", 2, "Number of infrastructure nodes")
	cmdFlags.app = flag.Uint("app-nodes", 3, "Number of application nodes")
//	cmdFlags.containerized = flag.Bool("containerized", false, "Run openshift containerized")
	cmdFlags.terraform = flag.String("terraform", "../terraform", "Directory of terraform files")
}

func ParseFlags() {
	initFlags()
	flag.Parse()
	ActiveSettings = getSettings()
}

func getSettings() Settings {
	settings := Settings{}

	settings.NumMasters = *cmdFlags.masters
	settings.NumInfra = *cmdFlags.infra
	settings.NumApplications = *cmdFlags.app
	settings.TerraformDir = *cmdFlags.terraform

	return settings
}