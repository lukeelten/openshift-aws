resource "aws_lb" "internal-lb" {
  name = "api-internal-lb"
  load_balancer_type = "network"
  internal = true

  subnets = ["${aws_subnet.subnet-private-1.id}"]

  tags {
    Name = "${var.project} - Internal API Load Balancer"
    Project = "${var.project}"
    Type = "internal"
  }
}

resource "aws_lb_target_group" "internal-lb-tg1" {
  name     = "api-internal-lb-tg1"
  port     = 8443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}

resource "aws_lb_listener" "internal-lb-listener" {
  load_balancer_arn = "${aws_lb.internal-lb.arn}"
  port              = "8443"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.internal-lb-tg1.arn}"
    type             = "forward"
  }
}
