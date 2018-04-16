package openshift

import "aws"

func GenerateConfig(sshConfig string) *InventoryConfig {
	masters := aws.MasterNodes()
	infra := aws.InfraNodes()
	apps := aws.AppNodes()

	config := InventoryConfig{
		Debug: true,
		OriginRelease: "v3.7.1",
		RoutesDomain: "apps.cc-openshift.de",
		InternalMaster: "internal-api.cc-openshift.de",
		ExternalMaster: "master.cc-openshift.de",
		SshConfig: sshConfig,

		Masters: make([]Node, len(masters)),
		Infras: make([]Node, len(infra)),
		Apps: make([]Node, len(apps)),
	}

	for i, node := range masters {
		config.Masters[i] = convertNodeObject(node, false, false)
	}

	for i, node := range infra {
		config.Infras[i] = convertNodeObject(node, true, true)
	}

	for i, node := range apps {
		config.Apps[i] = convertNodeObject(node, false, true)
	}

	return &config
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