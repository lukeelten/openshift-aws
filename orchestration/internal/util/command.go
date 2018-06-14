package util

import (
	"os/exec"
	"os"
	"io"
)

type Command interface {
	Run() error
	RunDir(string) error
	RunDirWithArgs(string, ...string) error
	RunWithArgs(...string) error
	SetArgs(...string)
	AppendArg(string)
}

type CommandImpl struct {
	Command

	cmd string
	dir string
	args []string

	stdin io.Reader
	stdout io.Writer
	stderr io.Writer
}

func NewCommand(command string, args ...string) Command {
	return NewCommandDir("", command, args...)
}

func NewCommandDir(dir string, command string, args ...string) Command {
	cmd := CommandImpl{
		dir: dir,
		cmd: command,
		args: args,
		stdin: os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}

	return cmd
}

func (cmd CommandImpl) Run() error {
	command := cmd.prepareCommand()
	return command.Run()
}

func (cmd CommandImpl) RunDir(dir string) error {
	command := cmd.prepareCommand()
	command.Dir = dir
	return command.Run()
}

func (cmd CommandImpl) RunWithArgs(args ...string) error {
	argsCopy := make([]string, len(cmd.args))
	copy(argsCopy, cmd.args)

	cmdCopy := cmd
	cmdCopy.args = append(argsCopy, args...)
	return cmdCopy.Run()
}

func (cmd CommandImpl) prepareCommand() *exec.Cmd {
	command := exec.Command(cmd.cmd, cmd.args...)
	command.Stdin = cmd.stdin
	command.Stderr = cmd.stderr
	command.Stdout = cmd.stdout
	command.Dir = cmd.dir

	return command
}

func (cmd CommandImpl) RunDirWithArgs(dir string, args ...string) error {
	argsCopy := make([]string, len(cmd.args))
	copy(argsCopy, cmd.args)

	cmdCopy := cmd
	cmdCopy.args = append(argsCopy, args...)
	return cmdCopy.RunDir(dir)
}