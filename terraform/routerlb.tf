
resource "aws_lb" "router-lb" {
  name = "router-lb"
  load_balancer_type = "network"
  //security_groups = ["${aws_security_group.router-elb-sg.id}", "${aws_security_group.allow-all-sg.id}"]

  subnets = ["${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-2.id}", "${aws_subnet.subnet-public-3.id}"]

  tags {
    Name = "${var.project} - Router Load Balancer"
    Project = "${var.project}"
  }
}

resource "aws_lb_target_group" "router-lb-tg1" {
  name     = "router-lb-tg1"
  port     = 80
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}

resource "aws_lb_target_group_attachment" "infra1-to-lb" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg1.arn}"
  target_id        = "${module.infrastructure_node1.node-id}"
  port             = 80
}

resource "aws_lb_target_group_attachment" "infra2-to-lb" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg1.arn}"
  target_id        = "${module.infrastructure_node2.node-id}"
  port             = 80
}

resource "aws_lb_target_group_attachment" "infra3-to-lb" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg1.arn}"
  target_id        = "${module.infrastructure_node3.node-id}"
  port             = 80
}

resource "aws_lb_listener" "router-lb-listener" {
  load_balancer_arn = "${aws_lb.router-lb.arn}"
  port              = "80"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.router-lb-tg1.arn}"
    type             = "forward"
  }
}