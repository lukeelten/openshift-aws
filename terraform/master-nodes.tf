resource "aws_instance" "master-node1" {
  depends_on      = ["aws_nat_gateway.private-nat", "aws_route.private_route"]

  ami             = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["master"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"

  vpc_security_group_ids = ["${aws_security_group.master-sg.id}", "${aws_security_group.etcd-sg.id}", "${aws_security_group.allow-internal.id}"]
  subnet_id = "${aws_subnet.subnet-private-1.id}"

  root_block_device {
    volume_type = "gp2"
    volume_size = 50
  }

  lifecycle {
    create_before_destroy = true
  }

  tags {
    Type = "master"
    Name = "${var.project} - Master Node 1"
    Project = "${var.project}"
  }
}

resource "aws_instance" "master-node2" {
  depends_on      = ["aws_nat_gateway.private-nat", "aws_route.private_route"]

  ami             = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["master"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"

  vpc_security_group_ids = ["${aws_security_group.master-sg.id}", "${aws_security_group.etcd-sg.id}"]
  subnet_id = "${aws_subnet.subnet-private-1.id}"

  root_block_device {
    volume_type = "gp2"
    volume_size = 50
  }

  lifecycle {
    create_before_destroy = true
  }

  tags {
    Type = "master"
    Name = "${var.project} - Master Node 2"
    Project = "${var.project}"
  }
}

resource "aws_lb_target_group_attachment" "master1-to-lb" {
  target_group_arn = "${aws_lb_target_group.master-lb-tg1.arn}"
  target_id        = "${aws_instance.master-node1.id}"
  port             = 8443
}

resource "aws_lb_target_group_attachment" "master2-to-lb" {
  target_group_arn = "${aws_lb_target_group.master-lb-tg1.arn}"
  target_id        = "${aws_instance.master-node2.id}"
  port             = 8443
}