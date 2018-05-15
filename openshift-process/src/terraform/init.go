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

func NewConfig(dir string, pubKey string, settings *configuration.InputVars) *Config {
	config := Config{}
	config.inited = false
	config.Dir = dir
	config.Vars = CreateConfig(settings, pubKey)

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
	config.inited = util.ExecuteDir(config.Dir, "terraform", "init") == nil

	return config.inited
}

func (config *Config) Apply() error {
	if !config.inited {
		panic("Please init terraform before apply")
	}

	return util.ExecuteDir(config.Dir, "terraform", "apply", "-auto-approve")
}

func (config *Config) Plan() error {
	if !config.inited {
		panic("Please init terraform before plan")
	}

	return util.ExecuteDir(config.Dir, "terraform", "plan")
}

func (config *Config) Validate() error {
	if !config.inited {
		panic("Please init terraform before validate")
	}

	return util.ExecuteDir(config.Dir, "terraform", "validate")
}

func (config *Config) Destroy() error {
	if !config.inited {
		panic("Please init terraform before destroy")
	}

	return util.ExecuteDir(config.Dir, "terraform", "destroy", "-auto-approve")
}