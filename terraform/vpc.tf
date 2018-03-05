resource "aws_vpc" "vpc" {
  cidr_block                       = "10.10.0.0/16"
  enable_dns_hostnames             = true
#  assign_generated_ipv6_cidr_block = true

  tags {
    Name = "${var.project} - VPC"
    Project = "${var.project}"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = "${aws_vpc.vpc.id}"

  tags {
    Name = "${var.project} - Internet Gateway"
    Project = "${var.project}"
  }
}

resource "aws_route_table" "public-rt" {
  vpc_id = "${aws_vpc.vpc.id}"

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.igw.id}"
  }

  tags {
    Name = "${var.project} - Public Route Table"
    Project = "${var.project}"
  }
}

resource "aws_route" "private-to-internet" {
  route_table_id            = "${aws_vpc.vpc.main_route_table_id}"
  destination_cidr_block    = "0.0.0.0/0"
  gateway_id                = "${aws_internet_gateway.igw.id}"
}