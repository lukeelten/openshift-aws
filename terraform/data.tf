
data "aws_availability_zones" "frankfurt" {}


data "aws_ami" "centos" {
  most_recent = true


  filter {
    name   = "name"
    values = ["CentOS Linux 7 x86_64 HVM EBS *"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["679593333241"] # CentOS official
}

data "aws_route53_zone" "existing-zone" {
  name = "${var.zone}"
  private_zone = false
}
