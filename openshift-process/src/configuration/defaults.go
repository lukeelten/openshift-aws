package configuration


const (
	DEFAULT_AMI = "${data.aws_ami.centos.id}"

	DEFAULT_TYPE_MASTER = "m4.large"
	DEFAULT_TYPE_INFRA = "t2.medium"
	DEFAULT_TYPE_APP = "t2.medium"

	DEFAULT_ROOT_MASTER = 50
	DEFAULT_ROOT_INFRA = 25
	DEFAULT_ROOT_APP = 25
)
