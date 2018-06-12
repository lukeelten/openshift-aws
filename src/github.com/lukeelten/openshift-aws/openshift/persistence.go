package openshift

import (
	"os"
	"text/template"
	"strings"
	"github.com/lukeelten/openshift-aws/aws"
	"github.com/lukeelten/openshift-aws/configuration"
)

type PersistenceConfig struct {
	ProjectId string

	EfsId string
	Region string

	EnableEfs bool

	EncryptEbs bool
	EnableEbs bool

	EbsDefault bool
	EfsDefault bool

	Zones string
}


func NewPersistenceConfig(settings *configuration.InputVars) *PersistenceConfig {
	config := PersistenceConfig{
		ProjectId: settings.ProjectId,
		EfsId:      "",
		Region:     settings.AwsConfig.Region,
		EnableEfs:  settings.Storage.EnableEfs,
		EncryptEbs: settings.Storage.EncryptEbs,
		EnableEbs:  settings.Storage.EnableEbs,
		EbsDefault: settings.Storage.Default == "ebs",
		EfsDefault: settings.Storage.Default == "efs",
	}

	return &config
}

func (config *PersistenceConfig) GeneratePersistenceConfigFiles(dir string) error {
	efsFilename := dir + "efs.yml"
	ebsFilename := dir + "ebs.yml"

	if config.EnableEfs {
		config.EfsId = aws.GetEFSId(config.ProjectId)

		f, err := os.Create(efsFilename)
		if err != nil {
			return err
		}
		defer f.Close()


		t, err := template.New("efs.tmpl").ParseFiles("templates/efs.tmpl")
		if err != nil {
			return err
		}

		err = t.Execute(f, config)

		if err != nil {
			return err
		}
	}

	if config.EnableEbs {
		zones := aws.GetAvailabilityZones()
		config.Zones = strings.Join(zones, ", ")

		f2, err := os.Create(ebsFilename)
		if err != nil {
			return err
		}
		defer f2.Close()


		t, err := template.New("ebs.tmpl").ParseFiles("templates/ebs.tmpl")
		if err != nil {
			return err
		}

		return t.Execute(f2, config)
	}

	return nil
}
