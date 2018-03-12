package ansible

import (
	"os/exec"
	"os"
	"io/ioutil"
)

type Playbook struct {
	filename string
}

func OpenPlaybook(filename string) *Playbook {
	playbook := &Playbook{filename}
	return playbook
}

func (playbook *Playbook) Run(inventory string) {
	cmd := exec.Command("ansible-playbook", "-i", inventory, playbook.filename)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ExecuteRemote (inventory string, nodes string, command string) {
	cmd := exec.Command("ansible", "-i", inventory, nodes, "-a", command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CheckReadiness (inventory string) bool {

	cmd := exec.Command("ansible", "-i", inventory, "nodes","-a", "/usr/bin/uname -a",  "-T", "5")
	cmd.Stdin = os.Stdin
	cmd.Stderr = ioutil.Discard
	cmd.Stdout = ioutil.Discard
	err := cmd.Run()

	return err != nil
}