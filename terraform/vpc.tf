resource "aws_vpc" "vpc" {
  cidr_block                       = "10.10.0.0/16"
  enable_dns_hostnames             = true
#  assign_generated_ipv6_cidr_block = true

  tags {
    Name = "${var.ProjectName} - VPC"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    "kubernetes.io/cluster/openshift" = "${var.ClusterId}"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = "${aws_vpc.vpc.id}"

  tags {
    Name = "${var.ProjectName} - Internet Gateway"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    "kubernetes.io/cluster/openshift" = "${var.ClusterId}"
  }
}

resource "aws_route_table" "public-rt" {
  vpc_id = "${aws_vpc.vpc.id}"

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.igw.id}"
  }

  tags {
    Name = "${var.ProjectName} - Public Route Table"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    "kubernetes.io/cluster/openshift" = "${var.ClusterId}"
  }
}

resource "aws_key_pair" "public-key" {
  key_name   = "${var.ProjectId}-key"
  public_key = "${var.PublicKey}"
}
