
resource "aws_elb" "internal-lb" {
  name     = "api-internal-lb"
  internal = true
  subnets = ["${aws_subnet.subnet-private-1.id}"]
  security_groups = ["${aws_security_group.allow-all-sg.id}"]

  listener {
    instance_port     = 8443
    instance_protocol = "tcp"
    lb_port           = 8443
    lb_protocol       = "tcp"
  }

  instances           = ["${aws_instance.master-node.*.id}"]

  health_check {
    healthy_threshold = 2
    interval = 5
    target = "TCP:8443"
    timeout = 4
    unhealthy_threshold = 5
  }

  tags {
    Name = "${var.project} - Internal Load Balancer"
    Project = "${var.project}"
    Type = "internal"
    ProjectId = "${var.project_id}"
  }
}