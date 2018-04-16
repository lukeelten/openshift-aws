package terraform

import (
	"os"
	"encoding/json"
	"regexp"
	"strings"
)

type AwsConfig struct {
	KeyId string
	SecretKey string
}

type Configuration struct {
	ProjectName string
	ProjectId string
	SshKey string
	Region string
	Zone string

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

func NewConfig() *Configuration {
	var config Configuration
	return &config
}

func DefaultConfig(ProjectName string, SshKey string, Zone string) *Configuration {
	config := NewConfig()

	config.ProjectName = ProjectName
	config.ProjectId = encodeProjectId(ProjectName)
	config.SshKey = SshKey
	config.Region = "eu-central-1"
	config.Zone = Zone

	config.Counts.Master = 2
	config.Counts.Infra = 2
	config.Counts.App = 2

	config.Types.Bastion = "t2.nano"
	config.Types.Master = "m4.large"
	config.Types.Infra = "t2.medium"
	config.Types.App = "t2.medium"

	return config
}

func (config* Configuration) GenerateJson() []byte {
	 b, err := json.Marshal(config)

	 if err != nil {
	 	panic(err)
	 }

	 return b
}

func (config* Configuration) WriteFile(filename string) {
	json := config.GenerateJson()

	f, err := os.Create(filename)
	if err != nil {
		panic(nil)
	}

	defer f.Close()

	f.Write(json)
	f.Sync()
}

func encodeProjectId(name string) string {
	r := regexp.MustCompile("[^\\w]")
	name = strings.ToLower(name)
	return r.ReplaceAllString(name, "")
}