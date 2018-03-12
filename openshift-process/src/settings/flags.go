package settings

import "flag"

type Settings struct {
	NumMasters uint
	NumInfra uint
	NumApplications uint
	Containerized bool
}

var ActiveSettings Settings

type flags struct {
	masters *uint
	infra *uint
	app *uint
	containerized *bool
}

var cmdFlags flags

func initFlags() {
	cmdFlags.masters = flag.Uint("master-nodes", 1, "Number of master nodes")
	cmdFlags.infra = flag.Uint("infra-nodes", 1, "Number of infrastructure nodes")
	cmdFlags.app = flag.Uint("app-nodes", 1, "Number of application nodes")
	cmdFlags.containerized = flag.Bool("containerized", false, "Run openshift containerized")

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
	settings.Containerized = *cmdFlags.containerized

	return settings
}