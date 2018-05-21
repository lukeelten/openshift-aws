package ansible

import (
	"util"
	"configuration"
)

type Playbook struct {
	filename string
}

var commands struct {
	ansible util.Command
	ansibleVerbose util.Command
	ansiblePlaybook util.Command
	ansiblePlaybookVerbose util.Command
}

func init() {
	commands.ansible = util.NewCommand("ansible", "-vvv", "-i")
	commands.ansibleVerbose = util.NewCommand("ansible", "-i")
	commands.ansiblePlaybook = util.NewCommand("ansible-playbook", "-i")
	commands.ansiblePlaybookVerbose = util.NewCommand("ansible-playbook", "-vvv", "-i")
}

func OpenPlaybook(filename string) *Playbook {
	playbook := &Playbook{filename}
	return playbook
}

func (playbook *Playbook) Run(inventory string) error {
	var cmd util.Command

	if configuration.Verbose {
		cmd = commands.ansiblePlaybookVerbose
	} else {
		cmd = commands.ansiblePlaybook
	}

	return cmd.RunWithArgs(inventory, playbook.filename)
}

func ExecuteRemote (inventory string, nodes string, command string) error {
	var cmd util.Command

	if configuration.Verbose {
		cmd = commands.ansibleVerbose
	} else {
		cmd = commands.ansible
	}

	return cmd.RunWithArgs(inventory, nodes, "-a", command)
}