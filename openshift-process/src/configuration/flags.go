package configuration

import (
	"flag"
	"util"
)

var Verbose bool

type CmdFlags struct {
	Debug bool

	ProjectId string
	ProjectName string

	AwsConfig struct {
		Region string
		KeyId string
		SecretKey string
	}

	ConfigFile string

	SkipTerraform bool
	SkipConfig    bool
	SkipPre       bool
	Verbose       bool
}


type flags struct {
	debug *bool

	projectName *string
	projectId *string

	region *string
	aws_key *string
	aws_secret *string

	configFile *string

	skipTerraform *bool
	existingConfig *bool
	skipPre *bool
	verbose *bool
}

var cmdFlags flags

func init() {
	cmdFlags.debug = flag.Bool("debug", true, "Debug mode disables some checks")
	cmdFlags.region = flag.String("region", "eu-central-1", "AWS region to create the infrastructure in")
	cmdFlags.aws_key = flag.String("aws-key", "", "AWS access key id. If empty the credentials used for AWS CLI will be loaded")
	cmdFlags.aws_secret = flag.String("aws-secret", "", "AWS secret key. If empty the credentials used for AWS CLI will be loaded")

	cmdFlags.projectName = flag.String("name", "", "Project Name to use in AWS tags and descriptions")
	cmdFlags.projectId = flag.String("id", "", "Project id to tag all instances. If empty an appropriate ID will be generated from project name.")
	cmdFlags.configFile = flag.String("config", "config.json", "Path / Name of configuration file to load")

	cmdFlags.skipTerraform = flag.Bool("skip-terraform", false, "Skip Terraform: Use when infrastructure already exist")
	cmdFlags.existingConfig = flag.Bool("skip-config", false, "Skip Config generation: Use when config already exist")
	cmdFlags.skipPre = flag.Bool("skip-pre", false, "Skip prerequisites playbook execution")
	cmdFlags.verbose = flag.Bool("verbose", false, "Verbose mode enables extended ansible output")
}

func ParseFlags() CmdFlags {
	flag.Parse()

	settings := CmdFlags{}
	loadValues(&settings)
	Verbose = settings.Verbose

	return settings
}

func loadValues(settings *CmdFlags) {
	settings.Debug = *cmdFlags.debug
	settings.ProjectName = *cmdFlags.projectName
	settings.ProjectId = *cmdFlags.projectId
	settings.ConfigFile = *cmdFlags.configFile
	settings.SkipTerraform = *cmdFlags.skipTerraform
	settings.SkipConfig = *cmdFlags.existingConfig
	settings.Verbose = *cmdFlags.verbose
	settings.SkipPre = *cmdFlags.skipPre

	if len(settings.ProjectName) >= NAME_MIN_LENGTH && len(settings.ProjectId) < NAME_MIN_LENGTH {
		settings.ProjectId = util.EncodeProjectId(settings.ProjectName)
	}

	settings.AwsConfig.Region = *cmdFlags.region
	settings.AwsConfig.KeyId = *cmdFlags.aws_key
	settings.AwsConfig.SecretKey = *cmdFlags.aws_secret
}