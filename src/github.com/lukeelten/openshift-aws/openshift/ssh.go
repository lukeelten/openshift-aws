package openshift

import (
	"os"
	"text/template"
	"github.com/lukeelten/openshift-aws/configuration"
	"github.com/lukeelten/openshift-aws/aws"
)

type SshConfig struct {
	BastionHostname string
}

func GenerateSshConfig(config *configuration.InputVars) *SshConfig {
	bastion := aws.BastionNode(config)
	sshConfig := SshConfig{bastion.ExternalDns}
	return &sshConfig
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