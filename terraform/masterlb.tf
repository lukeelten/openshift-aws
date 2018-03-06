/*
resource "aws_elb" "master-lb" {
  name               = "${var.project}-master-lb"

  listener {
    instance_port     = 443
    instance_protocol = "tcp"
    lb_port           = 443
    lb_protocol       = "tcp"
  }

  listener {
    instance_port      = 8443
    instance_protocol  = "tcp"
    lb_port            = 8443
    lb_protocol        = "tcp"
  }

  security_groups      = ["${aws_security_group.master-elb-sg.id}"]
  instances            = ["${module.master_node1.node-id}", "${module.master_node2.node-id}"]
  subnets = ["${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-2.id}", "${aws_subnet.subnet-public-3.id}"]
  depends_on = ["aws_internet_gateway.igw"]

  tags {
    Name = "${var.project} - Master Load Balancer"
    Project = "${var.project}"
  }
} */