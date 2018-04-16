package aws

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/aws"
)

type LoadBalancer struct {
	Dns string
}


func GetMasterLB() LoadBalancer {
	lb := getLbByName("master-lb")
	return LoadBalancer{*lb.DNSName}
}

func GetInfraLB() LoadBalancer {
	lb := getLbByName("router-lb")
	return LoadBalancer{*lb.DNSName}
}

func GetInternalLB() LoadBalancer {
	lb := getLbByName("api-internal-lb")
	return LoadBalancer{*lb.DNSName}
}

func getLbByName(name string) *elbv2.LoadBalancer {
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
