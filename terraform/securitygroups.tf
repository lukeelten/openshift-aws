resource "aws_security_group" "bastion-sg" {
  description = "${var.ProjectName} Security Group for Bastion server"
  name        = "${var.ProjectId}-bastion-sg"
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

resource "aws_security_group" "internal-lb-sg" {
  description = "${var.ProjectName} Security Group for Internal Master Load Balancer"
  name        = "${var.ProjectId}-internal-lb-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 8443
      to_port          = 8443
      protocol         = "tcp"
      security_groups  = ["${aws_security_group.nodes-sg.id}"]
    }
  ]

  egress {
    from_port        = 8443
    to_port          = 8443
    protocol         = "tcp"
    security_groups  = ["${aws_security_group.nodes-sg.id}"]
  }

  tags {
    Name = "${var.ProjectName} - Internal Load Balancer SG"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "master-sg" {
  description = "${var.ProjectName} Security Group for Master Nodes"
  name        = "${var.ProjectId}-master-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 8053
      to_port          = 8053
      protocol         = "tcp"
      security_groups  = ["${aws_security_group.nodes-sg.id}"]
    },
    {
      from_port        = 8053
      to_port          = 8053
      protocol         = "udp"
      security_groups  = ["${aws_security_group.nodes-sg.id}"]
    },
    {
      from_port        = 8443
      to_port          = 8443
      protocol         = "tcp"
      // Should be restricted to master load balancer, but network lbs does not have security groups
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      // Seems useless, but is for future use
      from_port        = 8443
      to_port          = 8443
      protocol         = "tcp"
      security_groups  = ["${aws_security_group.internal-lb-sg.id}", "${aws_security_group.nodes-sg.id}"]
    },
    {
      // Seems useless, but is for future use
      from_port        = 8444
      to_port          = 8444
      protocol         = "tcp"
      security_groups  = ["${aws_security_group.internal-lb-sg.id}", "${aws_security_group.nodes-sg.id}"]
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
  name        = "${var.ProjectId}-etcd-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 2379
      to_port          = 2379
      protocol         = "tcp"
      security_groups  = ["${aws_security_group.nodes-sg.id}"]
    },
    {
      from_port        = 2380
      to_port          = 2380
      protocol         = "tcp"
      self             = true
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
  name        = "${var.ProjectId}-infra-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 80
      to_port          = 80
      protocol         = "tcp"
      // Should be restricted to router lb
      cidr_blocks      = ["0.0.0.0/0"]
    },
    {
      from_port        = 443
      to_port          = 443
      protocol         = "tcp"
      // Should be restricted to router lb
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
    Name = "${var.ProjectName} - Infrastructure Nodes SG"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_security_group" "nodes-sg" {
  description = "${var.ProjectName} Security Group for Nodes"
  name        = "${var.ProjectId}-nodes-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 22
      to_port          = 22
      protocol         = "tcp"
      security_groups  = ["${aws_security_group.bastion-sg.id}"]
    },
    {
      from_port        = 10250
      to_port          = 10250
      protocol         = "tcp"
      self             = true
    },
    {
      from_port        = 10256
      to_port          = 10256
      protocol         = "tcp"
      self             = true
    },
    {
      from_port        = 4789
      to_port          = 4789
      protocol         = "udp"
      self             = true
    },
    {
      from_port        = 53
      to_port          = 53
      protocol         = "udp"
      self             = true
    },
    {
      from_port        = 53
      to_port          = 53
      protocol         = "tcp"
      self             = true
    },
    {
      // EFS
      from_port        = 2049
      to_port          = 2049
      protocol         = "tcp"
      security_groups  = ["${aws_security_group.storage-sg.id}"]
    },
    {
      // Elastic Search
      from_port        = 9300
      to_port          = 9300
      protocol         = "tcp"
      self             = true
    },
    {
      // Elastic Search
      from_port        = 9200
      to_port          = 9200
      protocol         = "tcp"
      self             = true
    },
    {
      // Fluentd
      from_port        = 9880
      to_port          = 9880
      protocol         = "tcp"
      self             = true
    },
    {
      // Fluentd
      from_port        = 24224
      to_port          = 24224
      protocol         = "tcp"
      self             = true
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
    "kubernetes.io/cluster/openshift" = "${var.ClusterId}"
  }
}

resource "aws_security_group" "storage-sg" {
  description = "${var.ProjectName} Storage Security Group"
  name        = "${var.ProjectId}-storage-sg"
  vpc_id      = "${aws_vpc.vpc.id}"

  ingress = [
    {
      from_port        = 2049
      to_port          = 2049
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
    {
      from_port        = 111
      to_port          = 111
      protocol         = "tcp"
      cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
    },
  ]

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["${aws_vpc.vpc.cidr_block}"]
  }

  tags {
    Name = "${var.ProjectName} - Storage Security Group"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}