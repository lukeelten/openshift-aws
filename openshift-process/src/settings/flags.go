package settings

import "flag"

type Settings struct {
	Debug bool
	ActivateTSB bool
}

var ActiveSettings Settings

type flags struct {
	debug *bool
	tsb *bool
}

var cmdFlags flags

func initFlags() {
	cmdFlags.debug = flag.Bool("debug", false, "Debug mode enables extended output")
	cmdFlags.tsb = flag.Bool("tsb", false, "Enable Template Service broker")
}

func ParseFlags() {
	initFlags()
	flag.Parse()
	ActiveSettings = getSettings()
}

func getSettings() Settings {
	settings := Settings{}
	settings.Debug = *cmdFlags.debug
	settings.ActivateTSB = *cmdFlags.tsb

	return settings
}