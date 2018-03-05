provider "aws" {
  access_key = "AKIAJ6SXYG5AOF6WER5Q"
  secret_key = "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"
  region     = "eu-central-1"
}

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
