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
}


func NewPersistenceConfig(settings *configuration.CmdFlags) *PersistenceConfig {
	efsId := aws.GetEFSId(settings.ProjectId)
	config := PersistenceConfig{efsId, settings.AWSConfig.Region}

	return &config
}

func (config *PersistenceConfig) GeneratePersistenceConfigFiles(dir string) error {
	deploymemtFilename := dir + "efs.yml"
	rolesFilename := dir + "efs-roles.yml"

	f, err := os.Create(deploymemtFilename)
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

	f2, err := os.Create(rolesFilename)
	if err != nil {
		return err
	}
	defer f2.Close()

	t2, err := template.New("efs-roles.tmpl").ParseFiles("templates/efs-roles.tmpl")
	if err != nil {
		return err
	}

	return t2.Execute(f2, config)
}