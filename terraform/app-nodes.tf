resource "aws_launch_configuration" "application-lc" {
  depends_on      = ["aws_internet_gateway.igw"]
  name            = "${var.ProjectName}-application-lc"
  image_id        = "${data.aws_ami.centos.id}"
  instance_type   = "${var.Types["App"]}"
  key_name        = "${var.SshKey}"
  user_data       = "${file("scripts/init.sh")}"
  security_groups = ["${aws_security_group.nodes-sg.id}", "${aws_security_group.allow-internal.id}"]

  root_block_device {
    volume_type = "gp2"
    volume_size = 25
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "application-scaling" {
  depends_on = ["aws_nat_gateway.private-nat", "aws_route.private_route"]
  name                 = "${var.ProjectName}-application-scaling-group"
  launch_configuration = "${aws_launch_configuration.application-lc.name}"

  min_size             = "${var.Counts["App"]}"
  max_size             = "${var.Counts["App"]}"

  vpc_zone_identifier  = ["${aws_subnet.subnet-private-1.id}"]

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
    value = "${var.ProjectName} - Application Node"
    propagate_at_launch = true
  }

  tag {
    key = "Project"
    value = "${var.ProjectName}"
    propagate_at_launch = true
  }

  tag {
    key = "ProjectId"
    value = "${var.ProjectId}"
    propagate_at_launch = true
  }
}
