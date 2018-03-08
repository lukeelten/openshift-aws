resource "aws_launch_configuration" "application-lc" {
  depends_on      = ["aws_internet_gateway.igw"]
  name            = "${var.project}-application-lc"
  image_id        = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["application"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"
  security_groups = ["${aws_security_group.nodes-sg.id}", "${aws_security_group.allow-all-sg.id}"]

  root_block_device {
    volume_type = "standard"
    volume_size = 25
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "application-scaling" {
  name                 = "${var.project}-application-scaling-group"
  launch_configuration = "${aws_launch_configuration.application-lc.name}"

  min_size             = 3
  max_size             = 3

  //  load_balancers       = ["${aws_elb.test-lb.id}"]
  vpc_zone_identifier  = ["${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-1.id}", "${aws_subnet.subnet-public-3.id}"]

  lifecycle {
    create_before_destroy = true
  }

  tag {
    key = "Type"
    value = "app"
    propagate_at_launch = true
  }

  tag {
    key = "Name"
    value = "${var.project} - Application Node"
    propagate_at_launch = true
  }

  tag {
    key = "Type"
    value = "${var.project}"
    propagate_at_launch = true
  }
}
