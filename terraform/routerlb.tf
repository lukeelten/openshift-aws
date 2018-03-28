
resource "aws_lb" "router-lb" {
  name = "router-lb"
  load_balancer_type = "network"

  subnets = ["${aws_subnet.subnet-public-1.id}"]

  tags {
    Name = "${var.project} - Router Load Balancer"
    Project = "${var.project}"
    Type = "infra"
  }
}

resource "aws_lb_listener" "router-lb-listener1" {
  load_balancer_arn = "${aws_lb.router-lb.arn}"
  port              = "80"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.router-lb-tg1.arn}"
    type             = "forward"
  }
}

resource "aws_lb_listener" "router-lb-listener2" {
  load_balancer_arn = "${aws_lb.router-lb.arn}"
  port              = "443"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.router-lb-tg2.arn}"
    type             = "forward"
  }
}

resource "aws_lb_target_group" "router-lb-tg1" {
  name     = "router-lb-tg1"
  port     = 80
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}

resource "aws_lb_target_group" "router-lb-tg2" {
  name     = "router-lb-tg2"
  port     = 443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}
