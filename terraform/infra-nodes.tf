resource "aws_launch_configuration" "infra-lc" {
  depends_on      = ["aws_internet_gateway.igw"]
  name            = "${var.project}-infra-lc"
  image_id        = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["infrastructure"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"
  security_groups = ["${aws_security_group.infra-sg.id}", "${aws_security_group.allow-all-sg.id}"]

  root_block_device {
    volume_type = "standard"
    volume_size = 25
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "infrastructure-scaling" {
  name                 = "${var.project}-infra-scaling-group"
  launch_configuration = "${aws_launch_configuration.infra-lc.name}"
  min_size             = 1
  max_size             = 1
  //  load_balancers       = ["${aws_elb.test-lb.id}"]
  vpc_zone_identifier  = ["${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-3.id}"]
  target_group_arns = ["${aws_lb_target_group.router-lb-tg1.arn}", "${aws_lb_target_group.router-lb-tg2.arn}"]

  lifecycle {
    create_before_destroy = true
  }

  tag {
    key = "Type"
    value = "infra"
    propagate_at_launch = true
  }

  tag {
    key = "Name"
    value = "${var.project} - Infrastructure Node"
    propagate_at_launch = true
  }

  tag {
    key = "Type"
    value = "${var.project}"
    propagate_at_launch = true
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
