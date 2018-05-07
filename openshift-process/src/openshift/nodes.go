package openshift

import (
	"aws"
	"configuration"
)

func GenerateConfig(sshConfig string, config *configuration.InputVars) *InventoryConfig {
	masters := aws.MasterNodes(config)
	infra := aws.InfraNodes(config)
	apps := aws.AppNodes(config)

	inventory := InventoryConfig{
		Debug: true,
		OriginRelease: "v3.7.2",
		RoutesDomain: "apps.cc-openshift.de",
		InternalMaster: "internal-api.cc-openshift.de",
		ExternalMaster: "master.cc-openshift.de",
		SshConfig: sshConfig,

		Masters: make([]Node, len(masters)),
		Infras: make([]Node, len(infra)),
		Apps: make([]Node, len(apps)),
	}

	for i, node := range masters {
		inventory.Masters[i] = convertNodeObject(node, false, false)
	}

	for i, node := range infra {
		inventory.Infras[i] = convertNodeObject(node, true, true)
	}

	for i, node := range apps {
		inventory.Apps[i] = convertNodeObject(node, false, true)
	}

	return &inventory
}

func convertNodeObject (nodeInfo aws.NodeInfo, infra bool, schedulable bool) Node {
	node := Node{
		InternalIp: nodeInfo.InternalIp,
		InternalHostname: nodeInfo.InternalDns,
		Zone: nodeInfo.Zone,
		Schedulable: schedulable,
		Region: "primary",
	}

	if infra {
		node.Region = "infra"
	}

	return node
}