package terraform

import (
	"settings"
	"os"
)

const (
	TYPE_MASTER = 1
	TYPE_INFRA = 2
	TYPE_APP = 3
)

type NodeInfo struct {
	internalIp string
	internalDns string
	externalIp string
	externalDns string

	terraformId string
	nodeType int
	loaded bool
}

type nodeConfig struct {
	name string

	ami string
	instanceType string
	key string
	vpc string

	userData string

	rootSize uint
}

var Nodes []NodeInfo

func GenerateNodes(filename string) {
	var fileContent string

	var i uint
	for i = 1; i <= settings.ActiveSettings.NumMasters; i++ {
		fileContent += generateMaster(i)
	}

	for i = 1; i <= settings.ActiveSettings.NumInfra; i++ {
		fileContent += generateInfra(i)
	}

	for i = 1; i <= settings.ActiveSettings.NumApplications; i++ {
		fileContent += generateApp(i)
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	f.WriteString(fileContent)
}


func generateMaster(cnt uint) string {
	id := "node-master-" + string(cnt)
	var config nodeConfig
	config.ami = settings.DEFAULT_AMI
	config.instanceType = settings.DEFAULT_TYPE_MASTER
	config.rootSize = settings.DEFAULT_ROOT_MASTER


	return generateNode(id, config)
}

func generateInfra(cnt uint) string {

}

func generateApp(cnt uint) string {

}

func generateNode(id string, config nodeConfig) string {


}