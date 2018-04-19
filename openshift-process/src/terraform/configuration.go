package terraform

import (
	"os"
	"encoding/json"
	"util"
)

type AwsConfig struct {
	KeyId string
	SecretKey string
}

type TerraformVars struct {
	ProjectName string
	ProjectId string
	SshKey string
	Zone string
	Region string

	Counts NodeCounts
	Types NodeTypes
}

type NodeCounts struct {
	Master uint
	Infra uint
	App uint
}

type NodeTypes struct {
	Bastion string
	Master string
	Infra string
	App string
}

func NewVars() *TerraformVars {
	var config TerraformVars
	return &config
}

func DefaultConfig(ProjectName string, SshKey string, Zone string) *TerraformVars {
	config := NewVars()

	config.ProjectName = ProjectName
	config.ProjectId = util.EncodeProjectId(ProjectName)
	config.SshKey = SshKey
	config.Zone = Zone
	config.Region = "eu-central-1"

	config.Counts.Master = 2
	config.Counts.Infra = 2
	config.Counts.App = 3

	config.Types.Bastion = "t2.nano"
	config.Types.Master = "m4.large"
	config.Types.Infra = "t2.medium"
	config.Types.App = "t2.medium"

	return config
}

func (config*TerraformVars) GenerateJson() []byte {
	 b, err := json.Marshal(config)

	 if err != nil {
	 	panic(err)
	 }

	 return b
}

func (config*TerraformVars) WriteFile(filename string) {
	json := config.GenerateJson()

	f, err := os.Create(filename)
	if err != nil {
		panic(nil)
	}

	defer f.Close()

	f.Write(json)
	f.Sync()
}

