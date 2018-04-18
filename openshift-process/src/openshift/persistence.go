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


func NewPersistenceConfig(settings *configuration.Settings) *PersistenceConfig {
	efsId := aws.GetEFSId(settings.ProjectId)
	config := PersistenceConfig{efsId, settings.AWSConfig.Region}

	return &config
}

func (config *PersistenceConfig) GeneratePersistenceConfigFiles(dir string) {
	deploymemtFilename := dir + "efs.yml"
	rolesFilename := dir + "efs-roles.yml"

	f, err := os.Create(deploymemtFilename)
	if err != nil {
		panic(nil)
	}
	defer f.Close()


	t, err := template.New("efs.tmpl").ParseFiles("templates/efs.tmpl")
	if err != nil {
		panic(err)
	}
	t.Execute(f, config)

	f2, err := os.Create(rolesFilename)
	if err != nil {
		panic(nil)
	}
	defer f.Close()

	t2, err := template.New("efs-roles.tmpl").ParseFiles("templates/efs-roles.tmpl")
	if err != nil {
		panic(err)
	}
	t2.Execute(f2, config)
}