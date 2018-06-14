
resource "aws_lb" "internal-lb" {
  depends_on      = ["aws_internet_gateway.igw"]
  name = "${var.ProjectId}-api-internal-lb"
  load_balancer_type = "network"

  subnets = ["${aws_subnet.subnets-private.*.id}"]
  internal = true

  count = "${var.Counts["Master"] > 1 ? 1 : 0}"

  tags {
    Name = "${var.ProjectName} - Internal Load Balancer"
    Project = "${var.ProjectName}"
    Type = "internal"
    ProjectId = "${var.ProjectId}"
    "kubernetes.io/cluster/openshift" = "${var.ClusterId}"
  }
}

resource "aws_lb_target_group" "internal-lb-tg1" {
  name     = "${var.ProjectId}-internal-lb-tg1"
  port     = 8443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"

  count = "${aws_lb.internal-lb.count}"

  tags {
    Name = "${var.ProjectName} - Traffic to Master Nodes"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }

  health_check {
    protocol = "TCP"
    interval = 10
    // timeout = 10
    // 30 seconds for a target to become healthy
    healthy_threshold = 3
    // 30 seconds to detect unhealthy targets
    unhealthy_threshold = 3
  }
}

resource "aws_lb_listener" "internal-lb-listener1" {
  count = "${aws_lb.internal-lb.count}"

  load_balancer_arn = "${aws_lb.internal-lb.arn}"
  port              = "8443"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.internal-lb-tg1.arn}"
    type             = "forward"
  }
}

resource "aws_lb_listener" "internal-lb-listener2" {
  count = "${aws_lb.internal-lb.count}"

  load_balancer_arn = "${aws_lb.internal-lb.arn}"
  port              = "443"
  protocol          = "TCP"

  default_action {
    target_group_arn = "${aws_lb_target_group.internal-lb-tg1.arn}"
    type             = "forward"
  }
}