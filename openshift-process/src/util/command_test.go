package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestCommandStdInOut(t *testing.T) {
	assert := assert.New(t)
	tmpCommand := NewCommand("test-command")
	cmd := tmpCommand.(CommandImpl)

	assert.Equal(os.Stdin, cmd.stdin)
	assert.Equal(os.Stdout, cmd.stdout)
	assert.Equal(os.Stderr, cmd.stderr)
}

func TestNewCommand(t *testing.T) {
	assert := assert.New(t)
	tmpCommand := NewCommand("test-command", "test1", "test2")
	cmd := tmpCommand.(CommandImpl)


	expectedArgs := []string{"test1", "test2"}

	assert.Equal("test-command", cmd.cmd)
	assert.Empty(cmd.dir)

	assert.Equal(2, len(cmd.args))
	assert.ElementsMatch(expectedArgs, cmd.args)

	tmpCommand = NewCommand("test-command2")
	cmd = tmpCommand.(CommandImpl)
	assert.Equal("test-command2", cmd.cmd)
	assert.Empty(cmd.dir)

	assert.Empty(cmd.args)
	assert.ElementsMatch(make([]string, 0), cmd.args)
}

func TestNewCommandDir(t *testing.T) {
	assert := assert.New(t)
	tmpCommand := NewCommandDir("test-dir", "test-command", "test1", "test2")
	cmd := tmpCommand.(CommandImpl)

	expectedArgs := []string{"test1", "test2"}

	assert.Equal("test-dir", cmd.dir)
	assert.Equal("test-command", cmd.cmd)

	assert.Equal(2, len(cmd.args))
	assert.ElementsMatch(expectedArgs, cmd.args)

	tmpCommand = NewCommandDir("test-dir", "test-command2")
	cmd = tmpCommand.(CommandImpl)

	assert.Equal("test-dir", cmd.dir)
	assert.Equal("test-command2", cmd.cmd)

	assert.Empty(cmd.args)
	assert.ElementsMatch(make([]string, 0), cmd.args)
}

func TestPrepareCommand(t *testing.T) {
	assert := assert.New(t)
	tmpCommand := NewCommandDir("test-dir", "test-command", "test1", "test2")
	command := tmpCommand.(CommandImpl)

	cmd := command.prepareCommand()
	assert.NotNil(cmd)

	assert.Equal(os.Stdin, cmd.Stdin)
	assert.Equal(os.Stdout, cmd.Stdout)
	assert.Equal(os.Stderr, cmd.Stderr)

	assert.Equal(command.dir, cmd.Dir)
	assert.Equal(command.cmd, cmd.Path)

	expectedArgs := []string{"test-command", "test1", "test2"}
	assert.ElementsMatch(expectedArgs, cmd.Args)
}