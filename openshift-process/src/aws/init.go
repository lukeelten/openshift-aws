package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/efs"
)

var Session *session.Session
var Client *ec2.EC2
var LBClent *elbv2.ELBV2
var EFSClient *efs.EFS

func (config *AwsConfig) InitSession() {
	awsConfig := aws.Config{
		Region: aws.String(config.Region),
	}

	if len(config.KeyId) > 0 && len(config.SecretKey) > 0 {
		awsConfig.Credentials = credentials.NewStaticCredentials(config.KeyId, config.SecretKey, "")
	}

	Session = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: awsConfig,
	}))

	Client = ec2.New(Session)
	LBClent = elbv2.New(Session)
	EFSClient = efs.New(Session)
}