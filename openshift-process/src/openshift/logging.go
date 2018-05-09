package openshift

import (
	"os"
	"text/template"
)



func (config *InventoryConfig) GenerateLoggingInventory(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	t, err := template.New("logging.tmpl").ParseFiles("templates/logging.tmpl")

	if err != nil {
		return err
	}

	return t.Execute(f, config)
}