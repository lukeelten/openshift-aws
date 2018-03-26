resource "aws_instance" "master-node" {
  depends_on      = ["aws_internet_gateway.igw"]
  name            = "${var.project}-master-lc"
  ami        = "${data.aws_ami.centos.id}"
  instance_type   = "${var.node-types["master"]}"
  key_name        = "${var.key}"
  user_data       = "${file("scripts/init.sh")}"
  security_groups = ["${aws_security_group.master-sg.id}", "${aws_security_group.etcd-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  subnet_id = "${aws_subnet.subnet-public-1}"

  root_block_device {
    volume_type = "gp2"
    volume_size = 50
  }

  lifecycle {
    create_before_destroy = true
  }

  tags {
    Type = "master"
    Name = "${var.project} - Master Node"
    Project = "${var.project}"
  }
}

/*
resource "aws_lb_target_group" "master-lb-tg1" {
  name     = "master-lb-tg1"
  port     = 8443
  protocol = "TCP"
  vpc_id   = "${aws_vpc.vpc.id}"
}
*/