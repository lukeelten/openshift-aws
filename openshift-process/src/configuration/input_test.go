package configuration

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func getValidInputVars() *InputVars {
	config := InputVars{
		ProjectName: "TestProject",
		ProjectId: "TestId",
		Domain: "test-cc.de",
		ClusterId: "1234",
		RegistryToS3: true,
		ClusterMetrics: true,
		AggregatedLogging: true,
		Debug: true,
	}

	config.NodeCounts.Master = 3
	config.NodeCounts.Infra = 3
	config.NodeCounts.App = 6

	config.Storage.EnableEfs = true
	config.Storage.EncryptEfs = true
	config.Storage.EnableEbs = true
	config.Storage.EncryptEbs = true

	config.NodeTypes.Bastion = "t2.nano"
	config.NodeTypes.Master = "m5.large"
	config.NodeTypes.Infra = "m5.large"
	config.NodeTypes.App = "m5.large"

	config.AwsConfig.Region = "eu-centra-1"
	config.AwsConfig.SecretKey = ""
	config.AwsConfig.KeyId = ""

	return &config
}

func TestDefaultConfig(t *testing.T) {
	assert := assert.New(t)

	config := DefaultConfig()
	assert.NotNil(config)
}

func TestInputVars_MergeCmdFlags(t *testing.T) {
	assert := assert.New(t)

	config := getValidInputVars()
	config.Debug = false
	assert.Nil(config.Validate())

	cmdFlags := CmdFlags{
		Debug: true,
		ProjectId: "cmd-project-id",
		ProjectName: "cmd-project-name",
	}
	cmdFlags.AwsConfig.Region = "eu-west-1"
	cmdFlags.AwsConfig.KeyId = "cmd-key-123"
	cmdFlags.AwsConfig.SecretKey = "#cmd-098#"

	assert.NotEqual("cmd-project-name", config.ProjectName)
	assert.NotEqual("cmd-project-id", config.ProjectId)

	assert.NotEqual("eu-west-1", config.AwsConfig.Region)
	assert.NotEqual("cmd-key-123", config.AwsConfig.KeyId)
	assert.NotEqual("#cmd-098#", config.AwsConfig.SecretKey)

	config.MergeCmdFlags(cmdFlags)
	assert.True(config.Debug)
	assert.Equal("cmd-project-name", config.ProjectName)
	assert.Equal("cmd-project-id", config.ProjectId)

	assert.Equal("eu-west-1", config.AwsConfig.Region)
	assert.Equal("cmd-key-123", config.AwsConfig.KeyId)
	assert.Equal("#cmd-098#", config.AwsConfig.SecretKey)

	config = getValidInputVars()
	config.Debug = true
	cmdFlags.Debug = false
	config.MergeCmdFlags(cmdFlags)
	assert.True(config.Debug)

	config = getValidInputVars()
	config.Debug = false
	cmdFlags.Debug = false
	config.MergeCmdFlags(cmdFlags)
	assert.False(config.Debug)

	config = getValidInputVars()
	oldKeyId := config.AwsConfig.KeyId
	oldSecretKey := config.AwsConfig.SecretKey
	cmdFlags.AwsConfig.KeyId = ""
	config.MergeCmdFlags(cmdFlags)

	assert.Equal(oldKeyId, config.AwsConfig.KeyId)
	assert.Equal(oldSecretKey, config.AwsConfig.SecretKey)

	config = getValidInputVars()
	oldKeyId = config.AwsConfig.KeyId
	oldSecretKey = config.AwsConfig.SecretKey
	cmdFlags.AwsConfig.SecretKey = ""
	config.MergeCmdFlags(cmdFlags)

	assert.Equal(oldKeyId, config.AwsConfig.KeyId)
	assert.Equal(oldSecretKey, config.AwsConfig.SecretKey)
}

func TestInputVars_Validate(t *testing.T) {
	assert := assert.New(t)

	// Init config
	config := getValidInputVars()
	assert.Nil(config.Validate())

	// Test Short Project Name
	config.ProjectName = "ab"
	assert.NotNil(config.Validate())
	// Test empty project name
	config.ProjectName = ""
	assert.NotNil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// Test Short project id
	config.ProjectId = "ab"
	assert.NotNil(config.Validate())

	// Test empty project id
	config.ProjectId = ""
	assert.NotNil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// Test empty cluster id
	config.ClusterId = ""
	assert.NotNil(config.Validate())

	testMasterCountValidation(assert)
	testInfraCountValidation(assert)
	testAppCountValidation(assert)

	testStorageValidation(assert)

	testDomainValidation(assert)

	testAwsValidation(assert)
}

func TestLoadInputVars(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	testFile := pwd + "/../../config.test.json"
	assert := assert.New(t)

	config := LoadInputVars(testFile)
	assert.NotNil(config)

	assert.Equal("TestProject", config.ProjectName)
	assert.Equal("TestId", config.ProjectId)
	assert.Equal("test.de", config.Domain)
	assert.Equal("1234", config.ClusterId)
	assert.True(config.AggregatedLogging)
	assert.False(config.ClusterMetrics)
	assert.True(config.RegistryToS3)

	assert.True(config.Storage.EnableEfs)
	assert.False(config.Storage.EncryptEfs)
	assert.False(config.Storage.EnableEbs)
	assert.True(config.Storage.EncryptEbs)
	assert.Equal("ebs", config.Storage.Default)

	assert.Equal(10, config.NodeCounts.Master)
	assert.Equal(1, config.NodeCounts.Infra)
	assert.Equal(0, config.NodeCounts.App)

	assert.Equal("t2.nano", config.NodeTypes.Bastion)
	assert.Equal("m5.xlarge", config.NodeTypes.Master)
	assert.Equal("m5.large", config.NodeTypes.Infra)
	assert.Equal("m5.xxlarge", config.NodeTypes.App)

	assert.Equal("eu-central-1", config.AwsConfig.Region)
	assert.Equal("key-123", config.AwsConfig.KeyId)
	assert.Equal("#9876#", config.AwsConfig.SecretKey)
}

func testMasterCountValidation(assert *assert.Assertions) {
	config := getValidInputVars()
	assert.Nil(config.Validate())

	// Test invalid master count
	config.NodeCounts.Master = 0
	assert.NotNil(config.Validate())
	config.NodeCounts.Master = -1
	assert.NotNil(config.Validate())
	config.NodeCounts.Master = -9999
	assert.NotNil(config.Validate())
	config.NodeCounts.Master = 2
	assert.NotNil(config.Validate())

	// Test valid master count
	config.NodeCounts.Master = 1
	assert.Nil(config.Validate())
	config.NodeCounts.Master = 3
	assert.Nil(config.Validate())
	config.NodeCounts.Master = 9999
	assert.Nil(config.Validate())
}

func testInfraCountValidation(assert *assert.Assertions) {
	config := getValidInputVars()
	assert.Nil(config.Validate())

	// Test invalid master count
	config.NodeCounts.Infra = 0
	assert.NotNil(config.Validate())
	config.NodeCounts.Infra = -1
	assert.NotNil(config.Validate())
	config.NodeCounts.Infra = -9999
	assert.NotNil(config.Validate())

	// Test valid master count
	config.NodeCounts.Infra = 1
	assert.Nil(config.Validate())
	config.NodeCounts.Infra = 3
	assert.Nil(config.Validate())
	config.NodeCounts.Infra = 9999
	assert.Nil(config.Validate())
}

func testAppCountValidation(assert *assert.Assertions) {
	config := getValidInputVars()
	assert.Nil(config.Validate())

	// Test invalid master count
	config.NodeCounts.App = 0
	assert.NotNil(config.Validate())
	config.NodeCounts.App = -1
	assert.NotNil(config.Validate())
	config.NodeCounts.App = -9999
	assert.NotNil(config.Validate())

	// Test valid master count
	config.NodeCounts.App = 1
	assert.Nil(config.Validate())
	config.NodeCounts.App = 3
	assert.Nil(config.Validate())
	config.NodeCounts.App = 9999
	assert.Nil(config.Validate())
}

func testStorageValidation(assert *assert.Assertions) {
	config := getValidInputVars()
	assert.Nil(config.Validate())

	// Test no active storage provider
	config.Storage.EnableEbs = false
	config.Storage.EnableEfs = false
	assert.NotNil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// test invalid default
	config.Storage.Default = "invalid"
	assert.NotNil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// test disabled default storage EBS
	config.Storage.Default = "ebs"
	config.Storage.EnableEbs = false
	assert.NotNil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// test disabled default storage EFS
	config.Storage.Default = "efs"
	config.Storage.EnableEfs = false
	assert.NotNil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// Aggregated logging require ebs
	config.AggregatedLogging = true
	config.Storage.EnableEbs = false
	assert.NotNil(config.Validate())
}


func testDomainValidation(assert *assert.Assertions) {
	config := getValidInputVars()
	assert.Nil(config.Validate())

	// test valid domain
	config.Domain = "codecentric.de"
	assert.Nil(config.Validate())

	// test valid subdomain
	config.Domain = "test.codecentric.de"
	assert.Nil(config.Validate())
	config.Domain = "sub.sub.sub.sub.codecentric.de"
	assert.Nil(config.Validate())

	// test invalid domain
	config.Domain = "abcd"
	assert.NotNil(config.Validate())
	config.Domain = "codecentric.codecentric11"
	assert.NotNil(config.Validate())

	// test invalid subdomain
	config.Domain = "123@dd.codecentric.de"
	assert.NotNil(config.Validate())
}

func testAwsValidation(assert *assert.Assertions) {
	config := getValidInputVars()
	assert.Nil(config.Validate())

	// test invalid region
	config.AwsConfig.Region = "europa-central-1"
	assert.NotNil(config.Validate())

	config.AwsConfig.Region = "eu-central"
	assert.NotNil(config.Validate())

	config.AwsConfig.Region = "eu-1"
	assert.NotNil(config.Validate())

	config.AwsConfig.Region = "eu-de-1"
	assert.NotNil(config.Validate())

	config.AwsConfig.Region = "eu-central-1a"
	assert.NotNil(config.Validate())

	// test valid region
	config.AwsConfig.Region = "eu-central-1"
	assert.Nil(config.Validate())

	config.AwsConfig.Region = "us-west-1"
	assert.Nil(config.Validate())

	config.AwsConfig.Region = "eu-west-1"
	assert.Nil(config.Validate())

	config.AwsConfig.Region = "ap-north-1"
	assert.Nil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// Test invalid node types bastion
	config.NodeTypes.Bastion = "t99.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Bastion = "z3.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Bastion = "t.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Bastion = "large"
	assert.NotNil(config.Validate())

	// test valid node types bastion
	config.NodeTypes.Bastion = "m4.xlarge"
	assert.Nil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// Test invalid node types master
	config.NodeTypes.Master = "t99.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Master = "z3.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Master = "t.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Master = "large"
	assert.NotNil(config.Validate())

	// test valid node types master
	config.NodeTypes.Master = "m4.xlarge"
	assert.Nil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// Test invalid node types infra
	config.NodeTypes.Infra = "t99.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Infra = "z3.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Infra = "t.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.Infra = "large"
	assert.NotNil(config.Validate())

	// test valid node types infra
	config.NodeTypes.Infra = "m4.xlarge"
	assert.Nil(config.Validate())

	// Reset config
	config = getValidInputVars()
	assert.Nil(config.Validate())

	// Test invalid node types app
	config.NodeTypes.App = "t99.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.App = "z3.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.App = "t.large"
	assert.NotNil(config.Validate())

	config.NodeTypes.App = "large"
	assert.NotNil(config.Validate())

	// test valid node types app
	config.NodeTypes.App = "m4.xlarge"
	assert.Nil(config.Validate())
}
