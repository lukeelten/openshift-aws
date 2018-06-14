package terraform

import (
	"../util"
	"../configuration"
)

type Config struct {
	inited bool

	Dir string
	State string
	VarsFile string
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

func NewConfig(dir string, state string, pubKey string, settings *configuration.InputVars) *Config {
	config := Config{}
	config.inited = false
	config.Dir = dir
	config.State = state
	config.Vars = CreateConfig(settings, pubKey)

	return &config
}

func (config *Config) GenerateVarsFile(filePath string) {
	config.VarsFile = filePath
	config.Vars.WriteFile(filePath)
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
	return commands.apply.RunDirWithArgs(config.Dir, config.getStateArgument(), config.getVarsFileArgument())
}

func (config *Config) Plan() error {
	config.checkState()
	return commands.plan.RunDirWithArgs(config.Dir, config.getStateArgument(), config.getVarsFileArgument())
}

func (config *Config) Validate() error {
	config.checkState()
	return commands.validate.RunDirWithArgs(config.Dir, config.getVarsFileArgument())
}

func (config *Config) Destroy() error {
	config.checkState()
	return commands.destroy.RunDirWithArgs(config.Dir, config.getStateArgument(), config.getVarsFileArgument())
}

func (config *Config) checkState() {
	if !config.inited {
		panic("Please init terraform before destroy")
	}
}

func (config *Config) getStateArgument() string {
	return "-state=" + config.State
}

func (config *Config) getVarsFileArgument() string {
	return "-var-file=" + config.VarsFile
}