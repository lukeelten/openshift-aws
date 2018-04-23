package configuration

import (
	"os"
	"util"
	"bufio"
	"io/ioutil"
	"encoding/json"
)

type InputVars struct {
	ProjectName string
	ProjectId string

	Domain string

	EnableEfs bool
	EncryptEfs bool

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

	vars.Validate()

	return &vars
}

func (vars *InputVars) Validate() bool {

}