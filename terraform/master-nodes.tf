resource "aws_launch_configuration" "master-lc" {
  depends_on      = ["aws_internet_gateway.igw"]
  name            = "${var.project}-master-lc"
  image_id        = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["master"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"
  security_groups = ["${aws_security_group.master-sg.id}", "${aws_security_group.etcd-sg.id}", "${aws_security_group.allow-all-sg.id}"]

  root_block_device {
    volume_type = "standard"
    volume_size = 50
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "master-scaling" {
  name                 = "${var.project}-master-scaling-group"
  launch_configuration = "${aws_launch_configuration.master-lc.name}"
  min_size             = 1
  max_size             = 1
  //  load_balancers       = ["${aws_elb.test-lb.id}"]
  vpc_zone_identifier  = ["${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-3.id}"]
  target_group_arns = ["${aws_lb_target_group.master-lb-tg1.arn}"]

  lifecycle {
    create_before_destroy = true
  }

  tag {
    key = "Type"
    value = "master"
    propagate_at_launch = true
  }

  tag {
    key = "Name"
    value = "${var.project} - Master Node"
    propagate_at_launch = true
  }

  tag {
    key = "Type"
    value = "${var.project}"
    propagate_at_launch = true
  }
}

resource "aws_lb_target_group" "master-lb-tg1" {
  name     = "master-lb-tg1"
  port     = 8443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}