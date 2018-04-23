package configuration

import (
	"flag"
	"aws"
	"util"
	"strings"
)

var Debug bool

type CmdFlags struct {
	Debug bool

	ProjectId string
	ProjectName string

	AWSConfig *aws.AwsConfig
}


type flags struct {
	debug *bool

	projectName *string
	projectId *string

	region *string
	aws_key *string
	aws_secret *string
}

var cmdFlags flags

func initFlags() {
	cmdFlags.debug = flag.Bool("debug", true, "Debug mode enables extended output")
	cmdFlags.region = flag.String("region", "eu-central-1", "AWS region to create the infrastructure in")
	cmdFlags.aws_key = flag.String("aws-key", "", "AWS access key id. If empty the credentials used for AWS CLI will be loaded")
	cmdFlags.aws_secret = flag.String("aws-secret", "", "AWS secret key. If empty the credentials used for AWS CLI will be loaded")

	cmdFlags.projectName = flag.String("name", "", "Project Name to use in AWS tags and descriptions")
	cmdFlags.projectId = flag.String("id", "", "Project id to tag all instances. If empty an appropriate ID will be generated from project name.")
}

func ParseFlags() CmdFlags {
	initFlags()
	flag.Parse()

	settings := CmdFlags{}
	loadValues(&settings)
	validateSettings(&settings)
	Debug = settings.Debug

	return settings
}

func loadValues(settings *CmdFlags) {
	settings.Debug = *cmdFlags.debug
	settings.ProjectName = *cmdFlags.projectName
	settings.ProjectId = *cmdFlags.projectId

	config := aws.NewConfig(*cmdFlags.region, *cmdFlags.aws_key, *cmdFlags.aws_secret)
	settings.AWSConfig = config
}

func validateSettings(settings *CmdFlags) {
	if len(settings.ProjectName) < 4 {
		panic("Invalid project name. Please provide at least 4 characters")
	}

	if len(settings.ProjectId) < 3 {
		settings.ProjectId = util.EncodeProjectId(settings.ProjectName)
	} else {
		id := util.EncodeProjectId(settings.ProjectId)
		if !strings.EqualFold(settings.ProjectId, id) {
			panic("The given project id contains invalid chracters. Please use only alphanumerical characters.")
		}
	}
}