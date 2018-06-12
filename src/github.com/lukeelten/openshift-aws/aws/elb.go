package aws

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/lukeelten/openshift-aws/configuration"
)

type LoadBalancer struct {
	Dns string
}


func GetMasterLB(config *configuration.InputVars) LoadBalancer {
	lb := getLbByName("master-lb", config)
	return LoadBalancer{*lb.DNSName}
}

func GetInfraLB(config *configuration.InputVars) LoadBalancer {
	lb := getLbByName("router-lb", config)
	return LoadBalancer{*lb.DNSName}
}

func GetInternalLB(config *configuration.InputVars) LoadBalancer {
	lb := getLbByName("api-internal-lb", config)
	return LoadBalancer{*lb.DNSName}
}

func getLbByName(name string, config *configuration.InputVars) *elbv2.LoadBalancer {
	name = config.ProjectId + "-" + name

	filter := &elbv2.DescribeLoadBalancersInput{
		Names: []*string{aws.String(name)},
	}

	result, err := LBClent.DescribeLoadBalancers(filter)

	if err != nil {
		panic(err)
	}

	if len(result.LoadBalancers) > 0 {
		return result.LoadBalancers[0]
	}

	return nil
}
