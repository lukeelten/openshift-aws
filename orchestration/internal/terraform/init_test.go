package terraform

import (
	"testing"
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"errors"
	"../configuration"
	"../util"
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

func TestConfig_CheckState(t *testing.T) {
	assert := assert2.New(t)

	config := &Config{}

	assert.False(config.inited)
	assert.Panics(func() {config.checkState()})

	config.inited = true
	assert.True(config.inited)
	assert.NotPanics(func() {config.checkState()})
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
	err := errors.New("test123")
	mock = commandMock{}
	mock.On("RunDir", testdir).Return(err)
	commands.init = mock

	ret = config.InitTerraform()

	mock.AssertExpectations(t)
	assert.NotNil(ret)
	assert.Equal(err, ret)
	assert.False(config.inited)
}

func TestConfig_Apply(t *testing.T) {
	const dir = "test-dir"
	resetCommands()

	mock := commandMock{}
	mock.On("RunDir", dir).Return(nil)
	commands.apply = mock

	assert := assert2.New(t)
	config := &Config{Dir:dir}

	assert.False(config.inited)
	assert.Panics(func() {config.Apply()})

	config.inited = true
	ret := config.Apply()
	assert.Nil(ret)
	mock.AssertExpectations(t)

	err := errors.New("test123")
	mock = commandMock{}
	mock.On("RunDir", dir).Return(err)
	commands.apply = mock

	ret = config.Apply()
	assert.NotNil(ret)
	assert.Equal(err, ret)
	mock.AssertExpectations(t)
}

func TestConfig_Destroy(t *testing.T) {
	const dir = "test-dir"
	resetCommands()

	mock := commandMock{}
	mock.On("RunDir", dir).Return(nil)
	commands.destroy = mock

	assert := assert2.New(t)
	config := &Config{Dir:dir}

	assert.False(config.inited)
	assert.Panics(func() {config.Apply()})

	config.inited = true
	ret := config.Destroy()
	assert.Nil(ret)
	mock.AssertExpectations(t)

	err := errors.New("test123")
	mock = commandMock{}
	mock.On("RunDir", dir).Return(err)
	commands.destroy = mock

	ret = config.Destroy()
	assert.NotNil(ret)
	assert.Equal(err, ret)
	mock.AssertExpectations(t)
}

func TestConfig_Plan(t *testing.T) {
	const dir = "test-dir"
	resetCommands()

	mock := commandMock{}
	mock.On("RunDir", dir).Return(nil)
	commands.plan = mock

	assert := assert2.New(t)
	config := &Config{Dir:dir}

	assert.False(config.inited)
	assert.Panics(func() {config.Apply()})

	config.inited = true
	ret := config.Plan()
	assert.Nil(ret)
	mock.AssertExpectations(t)

	err := errors.New("test123")
	mock = commandMock{}
	mock.On("RunDir", dir).Return(err)
	commands.plan = mock

	ret = config.Plan()
	assert.NotNil(ret)
	assert.Equal(err, ret)
	mock.AssertExpectations(t)
}

func TestConfig_Validate(t *testing.T) {
	const dir = "test-dir"
	resetCommands()

	mock := commandMock{}
	mock.On("RunDir", dir).Return(nil)
	commands.validate = mock

	assert := assert2.New(t)
	config := &Config{Dir:dir}

	assert.False(config.inited)
	assert.Panics(func() {config.Apply()})

	config.inited = true
	ret := config.Validate()
	assert.Nil(ret)
	mock.AssertExpectations(t)

	err := errors.New("test123")
	mock = commandMock{}
	mock.On("RunDir", dir).Return(err)
	commands.validate = mock

	ret = config.Validate()
	assert.NotNil(ret)
	assert.Equal(err, ret)
	mock.AssertExpectations(t)
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