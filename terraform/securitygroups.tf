resource "aws_security_group" "master-elb-sg" {
  description = "${var.project} Security Group for Master Load Balancer"
  name        = "${var.project}-master-elb-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 443
      to_port          = 443
      protocol         = "tcp"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = 8443
      to_port          = 8443
      protocol         = "tcp"
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = -1
      to_port          = -1
      protocol         = "icmp"
      cidr_blocks      = ["0.0.0.0/0"]
    }
  ]
  # Maybe egress restrict to ICMP and internal cidr
  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.project} - Master Load Balancer SG"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "internal-elb-sg" {
  description = "${var.project} Security Group for Internal Load Balancer"
  name        = "${var.project}-internal-elb-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 443
      to_port          = 443
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    }
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
  }

  tags {
    Name = "${var.project} - Internal Load Balancer SG"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "bastion-sg" {
  description = "${var.project} Security Group for Bastion server"
  name        = "${var.project}-bastion-sg"
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
    Name = "${var.project} - Bastion SG"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "master-sg" {
  description = "${var.project} Security Group for Master Nodes"
  name        = "${var.project}-master-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 443
      to_port          = 443
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
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
    Name = "${var.project} - Master Nodes SG"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "etcd-sg" {
  description = "${var.project} Security Group for ETCD"
  name        = "${var.project}-etcd-sg"
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
    Name = "${var.project} - ETCD SG"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "router-elb-sg" {
  description = "${var.project} Security Group for Router Load Balancer"
  name        = "${var.project}-router-elb-sg"
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
    cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
  }

  tags {
    Name = "${var.project} - Router Load Balancer SG"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "infra-sg" {
  description = "${var.project} Security Group for Infrastructure Nodes"
  name        = "${var.project}-infra-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 80
      to_port          = 80
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = 443
      to_port          = 443
      protocol         = "tcp"
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
    Name = "${var.project} - Infrastructure Nodes SG"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "nodes-sg" {
  description = "${var.project} Security Group for Nodes"
  name        = "${var.project}-nodes-sg"
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
    Name = "${var.project} - Nodes SG"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "allow-all-sg" {
  description = "${var.project} Allow everything"
  name        = "${var.project}-allow-all-sg"
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
    Name = "${var.project} - Allow All"
    Project = "${var.project}"
  }
}

resource "aws_security_group" "allow-internal" {
  description = "${var.project} Allow Internal Traffic"
  name        = "${var.project}-allow-internal-sg"
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
    }
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
  }

  tags {
    Name = "${var.project} - Allow All"
    Project = "${var.project}"
  }
}