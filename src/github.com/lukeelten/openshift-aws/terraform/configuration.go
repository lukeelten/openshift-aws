package terraform

import (
	"os"
	"encoding/json"
	"github.com/lukeelten/openshift-aws/configuration"
	"github.com/lukeelten/openshift-aws/util"
)

type AwsConfig struct {
	KeyId string
	SecretKey string
}

type TerraformVars struct {
	ProjectName string
	ProjectId string
	PublicKey string
	Zone string
	Region string

	Counts NodeCounts
	Types NodeTypes

	EnableEfs bool
	EncryptEfs bool

	RegistryS3 bool

	ClusterId string
}

type NodeCounts struct {
	Master int
	Infra int
	App int
}

type NodeTypes struct {
	Bastion string
	Master string
	Infra string
	App string
}

func CreateConfig(vars *configuration.InputVars, publicKey string) *TerraformVars {
	config := TerraformVars{
		ProjectName: vars.ProjectName,
		ProjectId: vars.ProjectId,
		Zone: vars.Domain,
		PublicKey: publicKey,
		Region: vars.AwsConfig.Region,
		Counts: NodeCounts{
			vars.NodeCounts.Master,
			vars.NodeCounts.Infra,
			vars.NodeCounts.App,
		},
		Types: NodeTypes{
			vars.NodeTypes.Bastion,
			vars.NodeTypes.Master,
			vars.NodeTypes.Infra,
			vars.NodeTypes.App,
		},

		EnableEfs: vars.Storage.EnableEfs,
		EncryptEfs: vars.Storage.EncryptEfs,
		RegistryS3: vars.RegistryToS3,
		ClusterId: vars.ClusterId,
	}

	return &config
}

func (config *TerraformVars) GenerateJson() []byte {
	 b, err := json.Marshal(config)
	 util.ExitOnError("Cannot marshal terraform configuration to json", err)

	 return b
}

func (config *TerraformVars) WriteFile(filename string) {
	data := config.GenerateJson()

	f, err := os.Create(filename)
	util.ExitOnError("Cannot create terraform configuration file", err)
	defer f.Close()

	f.Write(data)
	f.Sync()
}

