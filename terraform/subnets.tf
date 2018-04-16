resource "aws_subnet" "subnet-public-1" {
  vpc_id            = "${aws_vpc.vpc.id}"
  availability_zone = "${data.aws_availability_zones.frankfurt.names[0]}"

  cidr_block              = "${cidrsubnet(aws_vpc.vpc.cidr_block, 8, 1)}"
  map_public_ip_on_launch = true

  tags {
    Name = "${var.ProjectName} - Public Subnet 1"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

/*
resource "aws_subnet" "subnet-public-2" {
  vpc_id            = "${aws_vpc.vpc.id}"
  availability_zone = "${data.aws_availability_zones.frankfurt.names[1]}"

  cidr_block              = "${cidrsubnet(aws_vpc.vpc.cidr_block, 8, 2)}"
  map_public_ip_on_launch = true

  tags {
    Name = "${var.project} - Public Subnet 2"
    Project = "${var.project}"
  }
}

resource "aws_subnet" "subnet-public-3" {
  vpc_id            = "${aws_vpc.vpc.id}"
  availability_zone = "${data.aws_availability_zones.frankfurt.names[length(data.aws_availability_zones.frankfurt.names) - 1]}"

  cidr_block              = "${cidrsubnet(aws_vpc.vpc.cidr_block, 8, 3)}"
  map_public_ip_on_launch = true

  tags {
    Name = "${var.project} - Public Subnet 3"
    Project = "${var.project}"
  }
}
*/

resource "aws_route_table_association" "public-1-to-rt" {
  subnet_id      = "${aws_subnet.subnet-public-1.id}"
  route_table_id = "${aws_route_table.public-rt.id}"
}

/*
resource "aws_route_table_association" "public-2-to-rt" {
  subnet_id      = "${aws_subnet.subnet-public-2.id}"
  route_table_id = "${aws_route_table.public-rt.id}"
}

resource "aws_route_table_association" "public-3-to-rt" {
  subnet_id      = "${aws_subnet.subnet-public-3.id}"
  route_table_id = "${aws_route_table.public-rt.id}"
}
*/

resource "aws_subnet" "subnet-private-1" {
  vpc_id            = "${aws_vpc.vpc.id}"
  availability_zone = "${data.aws_availability_zones.frankfurt.names[0]}"

  cidr_block              = "${cidrsubnet(aws_vpc.vpc.cidr_block, 8, 4)}"
  map_public_ip_on_launch = false

  tags {
    Name = "${var.ProjectName} - Private Subnet 1"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

/*
resource "aws_subnet" "subnet-private-2" {
  vpc_id            = "${aws_vpc.vpc.id}"
  availability_zone = "${data.aws_availability_zones.frankfurt.names[1]}"

  cidr_block              = "${cidrsubnet(aws_vpc.vpc.cidr_block, 8, 5)}"
  map_public_ip_on_launch = false

  tags {
    Name = "${var.project} - Private Subnet 2"
    Project = "${var.project}"
  }
}

resource "aws_subnet" "subnet-private-3" {
  vpc_id            = "${aws_vpc.vpc.id}"
  availability_zone = "${data.aws_availability_zones.frankfurt.names[length(data.aws_availability_zones.frankfurt.names) - 1]}"

  cidr_block              = "${cidrsubnet(aws_vpc.vpc.cidr_block, 8, 6)}"
  map_public_ip_on_launch = false

  tags {
    Name = "${var.project} - Private Subnet 3"
    Project = "${var.project}"
  }
}
*/

resource "aws_route_table_association" "private-1-to-rt" {
  subnet_id      = "${aws_subnet.subnet-private-1.id}"
  route_table_id = "${aws_vpc.vpc.main_route_table_id}"
}

/*
resource "aws_route_table_association" "private-2-to-rt" {
  subnet_id      = "${aws_subnet.subnet-private-2.id}"
  route_table_id = "${aws_vpc.vpc.main_route_table_id}"
}

resource "aws_route_table_association" "private-3-to-rt" {
  subnet_id      = "${aws_subnet.subnet-private-3.id}"
  route_table_id = "${aws_vpc.vpc.main_route_table_id}"
}
*/