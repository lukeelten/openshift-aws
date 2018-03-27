resource "aws_instance" "infra-node1" {
  depends_on      = ["aws_internet_gateway.igw", "aws_nat_gateway.private-nat", "aws_route.private_route"]

  ami = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["infrastructure"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"
  vpc_security_group_ids = ["${aws_security_group.infra-sg.id}", "${aws_security_group.allow-internal.id}"]
  subnet_id = "${aws_subnet.subnet-private-1.id}"

  root_block_device {
    volume_type = "gp2"
    volume_size = 25
  }

  lifecycle {
    create_before_destroy = true
  }

  tags {
    Type = "infra"
    Name = "${var.project} - Infrastructure Node 1"
    Project = "${var.project}"
  }
}

resource "aws_instance" "infra-node2" {
  depends_on      = ["aws_internet_gateway.igw", "aws_nat_gateway.private-nat", "aws_route.private_route"]

  ami = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["infrastructure"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"
  vpc_security_group_ids = ["${aws_security_group.infra-sg.id}"]
  subnet_id = "${aws_subnet.subnet-private-1.id}"

  root_block_device {
    volume_type = "gp2"
    volume_size = 25
  }

  lifecycle {
    create_before_destroy = true
  }

  tags {
    Type = "infra"
    Name = "${var.project} - Infrastructure Node 2"
    Project = "${var.project}"
  }
}

resource "aws_lb_target_group_attachment" "infra-node1-tg1" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg1.arn}"
  target_id        = "${aws_instance.infra-node1.id}"
  port             = "${aws_lb_target_group.router-lb-tg1.port}"
}


resource "aws_lb_target_group_attachment" "infra-node1-tg2" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg2.arn}"
  target_id        = "${aws_instance.infra-node1.id}"
  port             = "${aws_lb_target_group.router-lb-tg2.port}"
}

resource "aws_lb_target_group_attachment" "infra-node2-tg1" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg1.arn}"
  target_id        = "${aws_instance.infra-node2.id}"
  port             = "${aws_lb_target_group.router-lb-tg1.port}"
}


resource "aws_lb_target_group_attachment" "infra-node2-tg2" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg2.arn}"
  target_id        = "${aws_instance.infra-node2.id}"
  port             = "${aws_lb_target_group.router-lb-tg2.port}"
}