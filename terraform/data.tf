
data "aws_availability_zones" "frankfurt" {}


data "aws_ami" "centos" {
  most_recent = true

  filter {
    name   = "name"
    values = ["CentOS Linux 7 x86_64 HVM EBS ENA *"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["679593333241"] # CentOS official
}

data "aws_route53_zone" "existing-zone" {
  name = "${var.Zone}"
  private_zone = false
}

output "ami-id" {
  value = "${data.aws_ami.centos.id}"
}

output "ami-name" {
  value = "${data.aws_ami.centos.name}"
}