package openshift

import (
	"../configuration"
	"../aws"
)

func GenerateConfig(sshConfig string, config *configuration.InputVars) *InventoryConfig {
	masters := aws.MasterNodes(config)
	infra := aws.InfraNodes(config)
	apps := aws.AppNodes(config)

	inventory := InventoryConfig{
		Debug: true,
		OriginRelease: "v3.9",
		RoutesDomain: "apps.cc-openshift.de",
		InternalMaster: "internal-api.cc-openshift.de",
		ExternalMaster: "master.cc-openshift.de",
		SshConfig: sshConfig,
		EnableEbs: config.Storage.EnableEbs,
		ClusterId: "1",

		AggregatedLogging: config.AggregatedLogging,
		ClusterMetrics: config.ClusterMetrics,
		RegistryToS3: config.RegistryToS3,

		Masters: make([]Node, len(masters)),
		Infras: make([]Node, len(infra)),
		Apps: make([]Node, len(apps)),
	}

	if inventory.RegistryToS3 {
		inventory.Registry.BucketName = aws.GetRegistryBucketName(config)
		inventory.Registry.Region = config.AwsConfig.Region
	}

	for i, node := range masters {
		inventory.Masters[i] = convertNodeObject(node, "master")
	}

	for i, node := range infra {
		inventory.Infras[i] = convertNodeObject(node, "infra")
	}

	for i, node := range apps {
		inventory.Apps[i] = convertNodeObject(node, "primary")
	}

	return &inventory
}

func convertNodeObject (nodeInfo aws.NodeInfo, region string) Node {
	node := Node{
		InternalIp: nodeInfo.InternalIp,
		InternalHostname: nodeInfo.InternalDns,
		Zone: nodeInfo.Zone,
		Region: region,
	}

	return node
}