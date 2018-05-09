package openshift

import (
	"configuration"
	"aws"
	"os"
	"text/template"
)

type PersistenceConfig struct {
	EfsId string
	Region string

	EnableEfs bool

	EncryptEbs bool
	EnableEbs bool

	EbsDefault bool
	EfsDefault bool
}


func NewPersistenceConfig(settings *configuration.InputVars) *PersistenceConfig {
	efsId := aws.GetEFSId(settings.ProjectId)
	config := PersistenceConfig{
		EfsId:      efsId,
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
