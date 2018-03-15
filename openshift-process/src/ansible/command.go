package ansible

import (
	"util"
)

type Playbook struct {
	filename string
}

func OpenPlaybook(filename string) *Playbook {
	playbook := &Playbook{filename}
	return playbook
}

func (playbook *Playbook) Run(inventory string) {
	util.Execute("ansible-playbook", "-i", inventory, playbook.filename)
}

func ExecuteRemote (inventory string, nodes string, command string) {
	util.Execute("ansible", "-i", inventory, nodes, "-a", command)
}

func CheckReadiness (inventory string) bool {
	return util.Execute("ansible", "-i", inventory, "nodes","-a", "/usr/bin/uname -a",  "-T", "5")
}