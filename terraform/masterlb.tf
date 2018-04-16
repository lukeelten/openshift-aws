
resource "aws_lb" "master-lb" {
  name = "master-lb"
  load_balancer_type = "network"

  subnets = ["${aws_subnet.subnet-public-1.id}"]

  tags {
    Name = "${var.ProjectName} - Master Load Balancer"
    Project = "${var.ProjectName}"
    Type = "master"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_lb_target_group" "master-lb-tg1" {
  name     = "master-lb-tg1"
  port     = 8443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}

resource "aws_lb_target_group" "master-lb-tg2" {
  name     = "master-lb-tg2"
  port     = 443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}

resource "aws_lb_listener" "master-lb-listener1" {
  load_balancer_arn = "${aws_lb.master-lb.arn}"
  port              = "8443"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.master-lb-tg1.arn}"
    type             = "forward"
  }
}

resource "aws_lb_listener" "master-lb-listener2" {
  load_balancer_arn = "${aws_lb.master-lb.arn}"
  port              = "443"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.master-lb-tg1.arn}"
    type             = "forward"
  }
}