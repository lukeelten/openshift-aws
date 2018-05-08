package openshift

import (
	"os"
	"text/template"
)

type LoggingData struct {
	Debug bool
	FirstMasterIp string
	ClusterId string
	SshConfig string
	OriginRelease string
}

func (config *InventoryConfig) GenerateLoggingInventory(filename string) error {
	data := LoggingData{config.Debug, config.Masters[0].InternalIp}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	t, err := template.New("logging.tmpl").ParseFiles("templates/logging.tmpl")

	if err != nil {
		return err
	}

	return t.Execute(f, data)
}