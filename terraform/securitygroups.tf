resource "aws_security_group" "bastion-sg" {
  description = "${var.ProjectName} Security Group for Bastion server"
  name        = "${var.ProjectName}-bastion-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 22
      to_port          = 22
      protocol         = "tcp"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = "-1"
      to_port          = "-1"
      protocol         = "icmp"
      cidr_blocks      = ["0.0.0.0/0"]
    }
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.ProjectName} - Bastion SG"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "master-sg" {
  description = "${var.ProjectName} Security Group for Master Nodes"
  name        = "${var.ProjectName}-master-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 8443
      to_port          = 8443
      protocol         = "tcp"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = 8053
      to_port          = 8053
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = 8053
      to_port          = 8053
      protocol         = "udp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = 22
      to_port          = 22
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    }
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.ProjectName} - Master Nodes SG"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "etcd-sg" {
  description = "${var.ProjectName} Security Group for ETCD"
  name        = "${var.ProjectName}-etcd-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 2379
      to_port          = 2379
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = "-1"
      to_port          = "-1"
      protocol         = "icmp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = 22
      to_port          = 22
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    }
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.ProjectName} - ETCD SG"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "infra-sg" {
  description = "${var.ProjectName} Security Group for Infrastructure Nodes"
  name        = "${var.ProjectName}-infra-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 80
      to_port          = 80
      protocol         = "tcp"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = 443
      to_port          = 443
      protocol         = "tcp"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = 22
      to_port          = 22
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = 2049
      to_port          = 2049
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.ProjectName} - Infrastructure Nodes SG"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "nodes-sg" {
  description = "${var.ProjectName} Security Group for Nodes"
  name        = "${var.ProjectName}-nodes-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 22
      to_port          = 22
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = 10250
      to_port          = 10250
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = 4789
      to_port          = 4789
      protocol         = "udp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.ProjectName} - Nodes SG"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "allow-all-sg" {
  description = "${var.ProjectName} Allow everything"
  name        = "${var.ProjectName}-allow-all-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 0
      to_port          = 0
      protocol         = "-1"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = -1
      to_port          = -1
      protocol         = "icmp"
      cidr_blocks      = ["0.0.0.0/0"]
    }
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.ProjectName} - Allow All"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "allow-internal" {
  description = "${var.ProjectName} Allow Internal Traffic"
  name        = "${var.ProjectName}-allow-internal-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 0
      to_port          = 0
      protocol         = "-1"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = -1
      to_port          = -1
      protocol         = "icmp"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = 2049
      to_port          = 2049
      protocol         = "TCP"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = 111
      to_port          = 111
      protocol         = "TCP"
      cidr_blocks      = ["0.0.0.0/0"]
    },
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
  }

  tags {
    Name = "${var.ProjectName} - Allow Internal Traffic"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "storage-sg" {
  description = "${var.ProjectName} Storage Security Group"
  name        = "${var.ProjectName}-storage-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 2049
      to_port          = 2049
      protocol         = "TCP"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = 111
      to_port          = 111
      protocol         = "TCP"
      cidr_blocks      = ["0.0.0.0/0"]
    },
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.ProjectName} - Storage Security Group"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}