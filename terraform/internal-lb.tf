
resource "aws_elb" "internal-lb" {
  name     = "api-internal-lb"
  internal = true
  subnets = ["${aws_subnet.subnets-private.*.id}"]
  security_groups = ["${aws_security_group.allow-all-sg.id}"]

  listener {
    instance_port     = 8443
    instance_protocol = "tcp"
    lb_port           = 8443
    lb_protocol       = "tcp"
  }

  instances           = ["${aws_instance.master-node.*.id}"]

  health_check {
    target = "TCP:8443"

    interval = 10
    timeout = 5
    // 50 seconds for a target to become healthy
    healthy_threshold = 5
    // 30 seconds to detect unhealthy targets
    unhealthy_threshold = 3
  }

  tags {
    Name = "${var.ProjectName} - Internal Load Balancer"
    Project = "${var.ProjectName}"
    Type = "internal"
    ProjectId = "${var.ProjectId}"
  }
}