package settings

import "os"

type SshConfig struct {
	filename string

	hosts []HostConfig
}

type HostConfig struct {
	Hostname string

	vars map[string]string
}

func NewSshConfig(filename string) *SshConfig {
	config := SshConfig{filename: filename}
	return &config
}

func (config *SshConfig) AddHost(host HostConfig) *SshConfig {
	config.hosts = append(config.hosts, host)
	return config
}

func NewHostConfig (hostname string) HostConfig {
	host := HostConfig{hostname, make(map[string]string)}
	return host
}

func (config *SshConfig) ToString() string {
	var content string

	for _,host := range config.hosts {
		content += "Host " + host.Hostname + "\n"

		for key,val := range host.vars {
			content += "\t" + key + " " + val + "\n"
		}

		content += "\n"
	}

	return content
}

func (config *SshConfig) Write() {
	f, err := os.Create(config.filename)
	if err != nil {
		panic(nil)
	}
	defer f.Close()

	f.WriteString(config.ToString())
	f.Sync()
}

func (host HostConfig) AddVar (key string, value string) HostConfig {
	host.vars[key] = value
	return host
}