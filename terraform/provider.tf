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

/*data "aws_ami" "fedora" {
  most_recent = true

  filter {
    name   = "name"
    values = ["Fedora-Atomic-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["125523088429"] # Fedora official
  name_regex = "^Fedora-Atomic-\\d-\\d{8}\\..*x86_64-eu-central-1-HVM"
} */

# data "aws_ami" "fedora_atomic" {
#  image_id = "ami-3d1e7352"
# }
