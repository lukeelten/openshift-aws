package terraform

import (
	"../util"
	"../configuration"
)

type Config struct {
	inited bool

	Dir string
	Vars *TerraformVars
}


var commands struct {
	init util.Command
	validate util.Command
	plan util.Command
	apply util.Command
	destroy util.Command
}

func init() {
	commands.init = util.NewCommand("terraform", "init")
	commands.validate = util.NewCommand("terraform", "validate")
	commands.plan = util.NewCommand("terraform", "plan")
	commands.apply = util.NewCommand("terraform", "apply", "-auto-approve")
	commands.destroy = util.NewCommand("terraform", "destroy", "-auto-approve")
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

func (config *Config) InitTerraform() error {
	if config.inited {
		return nil
	}

	err := commands.init.RunDir(config.Dir)
	config.inited = err == nil

	return err
}

func (config *Config) Apply() error {
	config.checkState()
	return commands.apply.RunDir(config.Dir)
}

func (config *Config) Plan() error {
	config.checkState()
	return commands.plan.RunDir(config.Dir)
}

func (config *Config) Validate() error {
	config.checkState()
	return commands.validate.RunDir(config.Dir)
}

func (config *Config) Destroy() error {
	config.checkState()
	return commands.destroy.RunDir(config.Dir)
}

func (config *Config) checkState() {
	if !config.inited {
		panic("Please init terraform before destroy")
	}
}