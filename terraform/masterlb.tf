
/*
resource "aws_lb" "master-lb" {
  name = "master-lb"
  load_balancer_type = "network"
  //security_groups = ["${aws_security_group.master-elb-sg.id}", "${aws_security_group.allow-all-sg.id}"]

  subnets = ["${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-2.id}", "${aws_subnet.subnet-public-3.id}"]

  tags {
    Name = "${var.project} - Master Load Balancer"
    Project = "${var.project}"
  }
}

resource "aws_lb_target_group" "master-lb-tg1" {
  name     = "master-lb-tg1"
  port     = 8443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}

resource "aws_lb_target_group_attachment" "master1-to-lb" {
  target_group_arn = "${aws_lb_target_group.master-lb-tg1.arn}"
  target_id        = "${module.master_node1.node-id}"
  port             = 8443
}

resource "aws_lb_target_group_attachment" "master2-to-lb" {
  target_group_arn = "${aws_lb_target_group.master-lb-tg1.arn}"
  target_id        = "${module.master_node2.node-id}"
  port             = 8443
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
*/