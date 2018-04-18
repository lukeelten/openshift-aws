package terraform

import (
	"configuration"
	"util"
)

type Config struct {
	inited bool

	Dir string
	Vars *TerraformVars
}

func NewConfig(dir string, settings *configuration.Settings) *Config {
	config := Config{}
	config.inited = false
	config.Dir = dir
	config.Vars = DefaultConfig(settings.ProjectName, "tobias@Codecentric", "cc-openshift.de")
	config.Vars.ProjectId = settings.ProjectId

	return &config
}

func (config *Config) GenerateVarsFile() {
	config.Vars.WriteFile(config.Dir + "/configuration.auto.tfvars")
}

func (config *Config) InitTerraform() bool {
	if config.inited {
		return true
	}

	config.GenerateVarsFile()
	config.inited = util.ExecuteDir(config.Dir, "terraform", "init")

	return config.inited
}

func (config *Config) Apply() {
	if !config.inited {
		panic("Please init terraform before apply")
	}

	util.ExecuteDir(config.Dir, "terraform", "apply", "-auto-approve")
}
