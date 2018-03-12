
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

/*
data "aws_ami" "fedora_atomic" {
  most_recent = true

  filter {
    name   = "name"
    values = ["Fedora-Atomic-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name = "image-id"
    values = ["ami-3d1e7352"]
  }

  owners = ["125523088429"] # Fedora official
}


data "aws_ami" "fedora_cloud" {
  most_recent = true

  filter {
    name = "image-id"
    values = ["ami-5f7cf830"]
  }
}


data "aws_ami" "ubuntu_server" {
  most_recent = true

  filter {
    name = "image-id"
    values = ["ami-5055cd3f"]
  }
}
*/