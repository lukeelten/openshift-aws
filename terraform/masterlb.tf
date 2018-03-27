

resource "aws_lb" "master-lb" {
  name = "master-lb"
  load_balancer_type = "network"
 // security_groups = ["${aws_security_group.master-elb-sg.id}"]

  subnets = ["${aws_subnet.subnet-public-1.id}"]

  tags {
    Name = "${var.project} - Master Load Balancer"
    Project = "${var.project}"
    Type = "master"
  }
}

resource "aws_lb_target_group" "master-lb-tg1" {
  name     = "master-lb-tg1"
  port     = 8443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}

resource "aws_lb_listener" "master-lb-listener" {
  load_balancer_arn = "${aws_lb.master-lb.arn}"
  port              = "8443"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.master-lb-tg1.arn}"
    type             = "forward"
  }
}
