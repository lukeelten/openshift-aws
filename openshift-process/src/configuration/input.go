package configuration

import (
	"util"
	"io/ioutil"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

const NAME_MIN_LENGTH=4

type InputVars struct {
	Debug bool

	ProjectName string
	ProjectId string

	Domain string

	AggregatedLogging bool

	Storage struct {
		EnableEfs bool
		EncryptEfs bool

		EnableEbs bool
		EncryptEbs bool

		Default string
	}

	NodeCounts struct {
		Master uint
		Infra uint
		App uint
	}

	NodeTypes struct {
		Bastion string
		Master string
		Infra string
		App string
	}

	AwsConfig struct {
		Region string
		KeyId string
		SecretKey string
	}
}


func LoadInputVars(filename string) *InputVars {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		util.ExitOnError("Cannot Open configuration file", err)
	}

	vars := InputVars{}
	json.Unmarshal(content, &vars)

	if len(vars.ProjectName) >= NAME_MIN_LENGTH && len(vars.ProjectId) < NAME_MIN_LENGTH {
		vars.ProjectId = util.EncodeProjectId(vars.ProjectName)
	}

	return &vars
}

func DefaultConfig() *InputVars {
	config := InputVars{}
	return &config
}

func (vars *InputVars) Validate() error {
	if len(vars.ProjectName) < NAME_MIN_LENGTH {
		return errors.New("invalid argument: Your Project name should contain at least 4 characters")
	}

	if len(vars.ProjectId) < NAME_MIN_LENGTH {
		return errors.New("invalid argument: Your Project ID should contain at least 4 characters")
	}

	if vars.NodeCounts.Master < 1 {
		return errors.New("invalid argument: Master Node Count cannot be 0")
	}

	if vars.NodeCounts.Infra < 1 {
		return errors.New("invalid argument: Infrastructure Node Count cannot be 0")
	}

	if vars.NodeCounts.App < 1 {
		return errors.New("invalid argument: Application Node Count cannot be 0")
	}

	// @todo validate instance types more precise
	r := regexp.MustCompile("[tmcpxridgfh][0-9]\\.[\\w]+")
	if !r.MatchString(vars.NodeTypes.Bastion) {
		return errors.New("invalid argument: Invalid Bastion type (" + vars.NodeTypes.Bastion + ")")
	}

	if !r.MatchString(vars.NodeTypes.Master) {
		return errors.New("invalid argument: Invalid Master type (" + vars.NodeTypes.Master + ")")
	}

	if !r.MatchString(vars.NodeTypes.Infra) {
		return errors.New("invalid argument: Invalid Infrastructure type (" + vars.NodeTypes.Infra + ")")
	}

	if !r.MatchString(vars.NodeTypes.App) {
		return errors.New("invalid argument: Invalid Application type (" + vars.NodeTypes.App + ")")
	}

	r = regexp.MustCompile("^([a-zA-Z0-9-_]+\\.)*[a-zA-Z0-9][a-zA-Z0-9-_]+\\.[a-zA-Z]{2,11}$")
	vars.Domain = strings.ToLower(vars.Domain)
	if !r.MatchString(vars.Domain) {
		return errors.New("invalid argument: Invalid Domain given (" + vars.Domain + ")")
	}

	r = regexp.MustCompile("[a-z]{2}-[a-z]{4,}-[\\d]")
	vars.AwsConfig.Region = strings.ToLower(vars.AwsConfig.Region)
	if !r.MatchString(vars.AwsConfig.Region) {
		return errors.New("invalid argument: Invalid AWS region (" + vars.AwsConfig.Region + ")")
	}

	return nil
}

func (vars *InputVars) MergeCmdFlags(flags CmdFlags) {
	vars.Debug = vars.Debug || flags.Debug

	if len(flags.ProjectName) > NAME_MIN_LENGTH {
		vars.ProjectName = flags.ProjectName
	}

	if len(flags.ProjectId) > NAME_MIN_LENGTH {
		vars.ProjectId = flags.ProjectId
	}

	if len(flags.AwsConfig.Region) > 0 {
		vars.AwsConfig.Region = flags.AwsConfig.Region
	}

	if len(flags.AwsConfig.KeyId) > 0 && len(flags.AwsConfig.SecretKey) > 0 {
		vars.AwsConfig.KeyId = flags.AwsConfig.KeyId
		vars.AwsConfig.SecretKey = flags.AwsConfig.SecretKey
	}
}