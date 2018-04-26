package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"configuration"
)

var Session *session.Session
var Client *ec2.EC2
var LBClent *elbv2.ELBV2
var EFSClient *efs.EFS

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
}