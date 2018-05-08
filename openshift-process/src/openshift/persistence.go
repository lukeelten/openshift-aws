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


func NewPersistenceConfig(settings *configuration.InputVars) *PersistenceConfig {
	efsId := aws.GetEFSId(settings.ProjectId)
	config := PersistenceConfig{efsId, settings.AwsConfig.Region}

	return &config
}

func (config *PersistenceConfig) GeneratePersistenceConfigFiles(dir string) error {
	deploymemtFilename := dir + "efs.yml"

	f, err := os.Create(deploymemtFilename)
	if err != nil {
		return err
	}
	defer f.Close()


	t, err := template.New("efs.tmpl").ParseFiles("templates/efs.tmpl")
	if err != nil {
		return err
	}

	return t.Execute(f, config)
}