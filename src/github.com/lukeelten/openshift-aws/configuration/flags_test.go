package configuration

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"util"
)


func TestLoadValues (t *testing.T) {
	assert := assert.New(t)

	cmdFlags = getFlagsStruct()
	settings := CmdFlags{}

	assert.Empty(settings.ProjectName)
	assert.Empty(settings.ProjectId)
	assert.Empty(settings.AwsConfig.Region)
	assert.Empty(settings.AwsConfig.KeyId)
	assert.Empty(settings.AwsConfig.SecretKey)
	assert.False(settings.Debug)
	assert.False(settings.SkipPre)
	assert.False(settings.SkipTerraform)
	assert.False(settings.SkipConfig)
	assert.False(settings.Verbose)

	loadValues(&settings)

	assert.Equal("test-name", settings.ProjectName)
	assert.Equal("test-id", settings.ProjectId)
	assert.Equal("test-region", settings.AwsConfig.Region)
	assert.Equal("test-key", settings.AwsConfig.KeyId)
	assert.Equal("test-secret", settings.AwsConfig.SecretKey)
	assert.True(settings.Debug)
	assert.True(settings.SkipPre)
	assert.False(settings.SkipTerraform)
	assert.False(settings.SkipConfig)
	assert.True(settings.Verbose)

	cmdFlags = getFlagsStruct()
	cmdFlags.projectId = stringPtr("")
	settings = CmdFlags{}

	loadValues(&settings)

	assert.NotEmpty(settings.ProjectId)
	assert.Equal(util.EncodeProjectId(settings.ProjectName), settings.ProjectId)
}

func getFlagsStruct() flags {
	flags := flags{
		debug: boolPtr(true),

		projectName: stringPtr("test-name"),
		projectId: stringPtr("test-id"),
		region: stringPtr("test-region"),
		aws_key: stringPtr("test-key"),
		aws_secret: stringPtr("test-secret"),

		skipPre: boolPtr(true),
		skipTerraform: boolPtr(false),
		existingConfig: boolPtr(false),
		verbose: boolPtr(true),

		configFile: stringPtr(""),
	}

	return flags
}

func boolPtr(value bool) *bool {
	return &value
}

func stringPtr(value string) *string {
	return &value
}