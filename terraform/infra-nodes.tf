resource "aws_instance" "infra-node" {
  depends_on      = ["aws_internet_gateway.igw"]
  name            = "${var.project}-infra-lc"
  ami = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["infrastructure"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"
  security_groups = ["${aws_security_group.infra-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  subnet_id = "${aws_subnet.subnet-public-1}"

  root_block_device {
    volume_type = "gp2"
    volume_size = 25
  }

  lifecycle {
    create_before_destroy = true
  }

  tags {
    Type = "infra"
    Name = "${var.project} - Infrastructure Node"
    Project = "${var.project}"
  }
}

/*
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
*/