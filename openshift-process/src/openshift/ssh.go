package openshift

import (
	"aws"
	"os"
	"text/template"
)

type SshConfig struct {
	BastionHostname string
}

func GenerateSshConfig() *SshConfig {
	bastion := aws.BastionNode()
	config := SshConfig{bastion.ExternalDns}
	return &config
}

func (config *SshConfig) WriteConfig(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	t, err := template.New("ssh.tmpl").ParseFiles("templates/ssh.tmpl")
	if err != nil {
		return err
	}

	return t.Execute(f, config)
}