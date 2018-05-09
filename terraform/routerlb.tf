
resource "aws_lb" "router-lb" {
  depends_on      = ["aws_internet_gateway.igw"]
  name = "${var.ProjectId}-router-lb"
  load_balancer_type = "network"

  subnets = ["${aws_subnet.subnets-public.*.id}"]

  tags {
    Name = "${var.ProjectName} - Router Load Balancer"
    Project = "${var.ProjectId}"
    Type = "infra"
    ProjectId = "${var.ProjectId}"
    "kubernetes.io/cluster/openshift" = "${var.ClusterId}"
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
  name     = "${var.ProjectId}-router-lb-tg1"
  port     = 80
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"

  tags {
    Name = "${var.ProjectName} - HTTP Routing Traffic"
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

resource "aws_lb_target_group" "router-lb-tg2" {
  name     = "${var.ProjectId}-router-lb-tg2"
  port     = 443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"

  tags {
    Name = "${var.ProjectName} - HTTPS Routing Traffic"
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
