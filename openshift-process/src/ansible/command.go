package ansible

import (
	"util"
	"configuration"
)

type Playbook struct {
	filename string
}

func OpenPlaybook(filename string) *Playbook {
	playbook := &Playbook{filename}
	return playbook
}

func (playbook *Playbook) Run(inventory string) error {
	if configuration.Verbose {
		return util.Execute("ansible-playbook", "-vvv", "-i", inventory, playbook.filename)
	} else {
		return util.Execute("ansible-playbook", "-i", inventory, playbook.filename)
	}
}

func ExecuteRemote (inventory string, nodes string, command string) error {
	if configuration.Verbose {
		return util.Execute("ansible", "-vvv", "-i", inventory, nodes, "-a", command)
	} else {
		return util.Execute("ansible", "-i", inventory, nodes, "-a", command)
	}
}

func CheckReadiness (inventory string) error {
	return util.Execute("ansible", "-i", inventory, "nodes","-a", "/usr/bin/uname -a",  "-T", "5")
}