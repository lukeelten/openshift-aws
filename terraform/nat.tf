resource "aws_nat_gateway" "private-nat" {
  depends_on      = ["aws_internet_gateway.igw"]

  allocation_id = "${aws_eip.nat-eip.id}"
  subnet_id     = "${aws_subnet.subnets-public.*.id[1]}"

  tags {
    Name = "${var.ProjectName} - Private Subnet"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    "kubernetes.io/cluster/openshift" = "${var.ClusterId}"
  }
}

resource "aws_eip" "nat-eip" {
  vpc      = true
  depends_on = ["aws_internet_gateway.igw"]

  tags {
    Name ="${var.ProjectName} - NAT Internet IP"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_route" "private_route" {
  route_table_id  = "${aws_vpc.vpc.main_route_table_id}"
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id = "${aws_nat_gateway.private-nat.id}"
}