package terraform

import (
	"testing"
	assert2 "github.com/stretchr/testify/assert"
	"configuration"
	"github.com/stretchr/testify/mock"
	"util"
	"errors"
)


func TestNewConfig(t *testing.T) {
	assert := assert2.New(t)

	inputVars := configuration.InputVars{}
	dir := "test-dir"
	publicKey := "ssh-rsa AAA..."

	config := NewConfig(dir, publicKey, &inputVars)

	assert.False(config.inited)
	assert.Equal(dir, config.Dir)
	assert.NotNil(config.Vars)
	assert.Equal(publicKey, config.Vars.PublicKey)
}

func TestConfig_InitTerraformSuccess(t *testing.T) {
	resetCommands()

	const testdir = "test"

	mock := commandMock{}
	mock.On("RunDir", testdir).Return(nil)
	commands.init = mock

	assert := assert2.New(t)

	inputVars := configuration.InputVars{}
	publicKey := "ssh-rsa AAA..."

	config := NewConfig(testdir, publicKey, &inputVars)
	ret := config.InitTerraform()

	mock.AssertExpectations(t)
	assert.Nil(ret)
	assert.True(config.inited)

	// now there should be no more calls to command
	mock = commandMock{}
	commands.init = mock
	ret = config.InitTerraform()

	assert.True(config.inited)
	mock.AssertExpectations(t)
	assert.Nil(ret)
}

func TestConfig_InitTerraformFail(t *testing.T) {
	resetCommands()

	const testdir = "test"

	mock := commandMock{}
	mock.On("RunDir", testdir).Return(errors.New("anything"))
	commands.init = mock

	assert := assert2.New(t)

	inputVars := configuration.InputVars{}
	publicKey := "ssh-rsa AAA..."

	config := NewConfig(testdir, publicKey, &inputVars)
	ret := config.InitTerraform()

	mock.AssertExpectations(t)
	assert.NotNil(ret)
	assert.False(config.inited)

	// Next call should do the same
	mock = commandMock{}
	mock.On("RunDir", testdir).Return(errors.New("anything"))
	commands.init = mock

	ret = config.InitTerraform()

	mock.AssertExpectations(t)
	assert.NotNil(ret)
	assert.False(config.inited)
}


func resetCommands() {
	// Resets all commands to nil. If an unexpected call is done, tests wll fail due to nil object
	commands.init = nil
	commands.validate = nil
	commands.plan = nil
	commands.apply = nil
	commands.destroy = nil
}

type commandMock struct {
	mock.Mock
	util.Command
}

func (mock commandMock)	Run() error {
	args := mock.Called()
	return args.Error(0)
}

func (mock commandMock)	RunDir(dir string) error {
	args := mock.Called(dir)
	return args.Error(0)
}

func (mock commandMock) RunWithArgs(arguments ...string) error {
	args := mock.Called(arguments)
	return args.Error(0)
}