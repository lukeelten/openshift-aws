package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

var Session *session.Session
var Client *ec2.EC2

type NodeInfo struct {
	InternalIp  string
	InternalDns string
	ExternalIp  string
	ExternalDns string
}

func InitAws() {
	Session = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	Client = ec2.New(Session)
}

func MasterNodes() []NodeInfo {
	return loadNodesOfType("master")
}

func InfraNodes() []NodeInfo {
	return loadNodesOfType("infra")
}

func AppNodes() []NodeInfo {
	return loadNodesOfType("app")
}

func loadNodesOfType(type_ string) []NodeInfo {
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Type"),
				Values: []*string{aws.String(type_)},
			},
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{aws.String("running")},
			},
		},
	}

	result, err := Client.DescribeInstances(params)

	if err != nil {
		panic(err)
	}

	return extractNodes(result)
}

func extractNodes(result *ec2.DescribeInstancesOutput) []NodeInfo {
	numInstances := len(result.Reservations)
	var nodes []NodeInfo

	for i := 0; i < numInstances; i++ {
		current := result.Reservations[i]

		if len(current.Instances) < 1 || len(current.Instances[0].NetworkInterfaces) < 1 {
			continue
		}

		state := *current.Instances[0].State.Name
		if state != "running" {
			continue
		}

		for _,instance := range current.Instances {
			nodes = append(nodes, extractNodeInfo(instance))
		}
	}

	return nodes
}

func extractNodeInfo(current *ec2.Instance) NodeInfo {
	var node NodeInfo

	network := current.NetworkInterfaces[0]

	node.InternalIp = *network.PrivateIpAddress
	node.InternalDns = *network.PrivateDnsName
	node.ExternalIp = *network.Association.PublicIp
	node.ExternalDns = *network.Association.PublicDnsName

	return node
}