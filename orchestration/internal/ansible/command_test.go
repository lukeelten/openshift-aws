package ansible

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"errors"
	"testing"
	"../configuration"
	"../util"
)

func resetCommands() {
	commands.ansible = nil
	commands.ansibleVerbose = nil
	commands.ansiblePlaybook = nil
	commands.ansiblePlaybookVerbose = nil
}

func TestOpenPlaybook(t *testing.T) {
	const playbookName = "test-playbook"
	assert := assert.New(t)

	playbook := OpenPlaybook(playbookName)
	assert.NotNil(playbook)
	assert.Equal(playbookName, playbook.filename)
}

func TestPlaybook_Run(t *testing.T) {
	const playbookName = "test-playbook"
	const inventory = "test-inventory"
	configuration.Verbose = false
	assert := assert.New(t)

	resetCommands()
	mock := CommandMock{}
	mock.On("RunWithArgs", []string{inventory, playbookName}).Return(nil)
	commands.ansiblePlaybook = mock

	playbook := OpenPlaybook(playbookName)
	assert.NotNil(playbook)

	ret := playbook.Run(inventory)
	assert.Nil(ret)
	mock.AssertExpectations(t)

	err := errors.New("test123")
	mock = CommandMock{}
	mock.On("RunWithArgs", []string{inventory, playbookName}).Return(err)
	commands.ansiblePlaybook = mock

	ret = playbook.Run(inventory)
	assert.NotNil(ret)
	assert.Equal(err, ret)
	mock.AssertExpectations(t)
}

func TestPlaybook_RunVerbose(t *testing.T) {
	const playbookName = "test-playbook"
	const inventory = "test-inventory"
	configuration.Verbose = true
	assert := assert.New(t)

	resetCommands()
	mock := CommandMock{}
	mock.On("RunWithArgs", []string{inventory, playbookName}).Return(nil)
	commands.ansiblePlaybookVerbose = mock

	playbook := OpenPlaybook(playbookName)
	assert.NotNil(playbook)

	ret := playbook.Run(inventory)
	assert.Nil(ret)
	mock.AssertExpectations(t)

	err := errors.New("test123")
	mock = CommandMock{}
	mock.On("RunWithArgs", []string{inventory, playbookName}).Return(err)
	commands.ansiblePlaybookVerbose = mock

	ret = playbook.Run(inventory)
	assert.NotNil(ret)
	assert.Equal(err, ret)
	mock.AssertExpectations(t)
}

func TestExecuteRemote(t *testing.T) {
	resetCommands()

	const (
		inventory = "inventory"
		nodes = "nodes"
		command = "test-cmd"
	)

	configuration.Verbose = false
	assert := assert.New(t)

	mock := CommandMock{}
	mock.On("RunWithArgs", []string{inventory, nodes, "-a", command}).Return(nil)
	commands.ansible = mock

	ret := ExecuteRemote(inventory, nodes, command)
	assert.Nil(ret)
	mock.AssertExpectations(t)

	err := errors.New("test123")
	mock = CommandMock{}
	mock.On("RunWithArgs", []string{inventory, nodes, "-a", command}).Return(err)
	commands.ansible = mock

	ret = ExecuteRemote(inventory, nodes, command)
	assert.NotNil(ret)
	assert.Equal(err, ret)
	mock.AssertExpectations(t)
}

func TestExecuteRemoteVerbose(t *testing.T) {
	resetCommands()

	const (
		inventory = "inventory"
		nodes = "nodes"
		command = "test-cmd"
	)

	configuration.Verbose = true
	assert := assert.New(t)

	mock := CommandMock{}
	mock.On("RunWithArgs", []string{inventory, nodes, "-a", command}).Return(nil)
	commands.ansibleVerbose = mock

	ret := ExecuteRemote(inventory, nodes, command)
	assert.Nil(ret)
	mock.AssertExpectations(t)

	err := errors.New("test123")
	mock = CommandMock{}
	mock.On("RunWithArgs", []string{inventory, nodes, "-a", command}).Return(err)
	commands.ansibleVerbose = mock

	ret = ExecuteRemote(inventory, nodes, command)
	assert.NotNil(ret)
	assert.Equal(err, ret)
	mock.AssertExpectations(t)
}

type CommandMock struct {
	mock.Mock
	util.Command
}

func (mock CommandMock)	Run() error {
	args := mock.Called()
	return args.Error(0)
}

func (mock CommandMock)	RunDir(dir string) error {
	args := mock.Called(dir)
	return args.Error(0)
}

func (mock CommandMock) RunWithArgs(arguments ...string) error {
	args := mock.Called(arguments)
	return args.Error(0)
}