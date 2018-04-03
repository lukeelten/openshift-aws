package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

var Session *session.Session
var Client *ec2.EC2
var LBClent *elbv2.ELBV2

type NodeInfo struct {
	InternalIp  string
	InternalDns string
	ExternalIp  string
	ExternalDns string
	Zone string
}

func InitAws() {
	Session = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	Client = ec2.New(Session)
	LBClent = elbv2.New(Session)
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

func BastionNode() NodeInfo {
	bastion := loadNodesOfType("bastion")
	if len(bastion) < 1 {
		panic("No bastion host found")
	}

	return bastion[0]
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

	node.InternalIp = *current.PrivateIpAddress
	node.InternalDns = *current.PrivateDnsName

	if current.PublicIpAddress != nil {
		node.ExternalIp = *current.PublicIpAddress
	}

	if current.PublicDnsName != nil {
		node.ExternalDns = *current.PublicDnsName
	}

	if current.Placement != nil {
		node.Zone = *current.Placement.AvailabilityZone
	}

	return node
}