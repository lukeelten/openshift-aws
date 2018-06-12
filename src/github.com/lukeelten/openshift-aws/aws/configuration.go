package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lukeelten/openshift-aws/util"
	"github.com/lukeelten/openshift-aws/configuration"
)

var Session *session.Session
var Client *ec2.EC2
var LBClent *elbv2.ELBV2
var EFSClient *efs.EFS
var S3Client *s3.S3

func InitSession(config *configuration.InputVars) {
	awsConfig := aws.Config{
		Region: aws.String(config.AwsConfig.Region),
	}

	if len(config.AwsConfig.KeyId) > 0 && len(config.AwsConfig.SecretKey) > 0 {
		awsConfig.Credentials = credentials.NewStaticCredentials(config.AwsConfig.KeyId, config.AwsConfig.SecretKey, "")
	}

	Session = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: awsConfig,
	}))

	Client = ec2.New(Session)
	LBClent = elbv2.New(Session)
	EFSClient = efs.New(Session)
	S3Client = s3.New(Session)
}

func GetAvailabilityZones() []string {
	result, err := Client.DescribeAvailabilityZones(nil)
	util.ExitOnError("Cannot load availability zones", err)

	zones := make([]string, len(result.AvailabilityZones))
	for i, zone := range result.AvailabilityZones {
		zones[i] = *zone.ZoneName
	}

	return zones
}