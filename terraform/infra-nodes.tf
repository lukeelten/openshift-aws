resource "aws_instance" "infra-node" {
  depends_on      = ["aws_internet_gateway.igw", "aws_nat_gateway.private-nat", "aws_route.private_route"]

  ami = "${data.aws_ami.centos.id}"
  instance_type   = "${var.Types["Infra"]}"
  key_name        = "${aws_key_pair.public-key.key_name}"
  user_data       = "${file("assets/init.sh")}"
  vpc_security_group_ids = ["${aws_security_group.nodes-sg.id}", "${aws_security_group.infra-sg.id}"]
  subnet_id = "${aws_subnet.subnets-private.*.id[(count.index % aws_subnet.subnets-private.count)]}"
  iam_instance_profile = "${aws_iam_instance_profile.infra-profile.name}"

  count = "${var.Counts["Infra"]}"

  root_block_device {
    volume_type = "gp2"
    volume_size = 40
  }

  lifecycle {
    create_before_destroy = true
  }

  tags {
    Type = "infra"
    Name = "${var.ProjectName} - Infrastructure Node ${count.index + 1}"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    "kubernetes.io/cluster/openshift" = "${var.ClusterId}"
  }
}

resource "aws_lb_target_group_attachment" "infra-node-tg1" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg1.arn}"
  target_id        = "${aws_instance.infra-node.*.id[count.index]}"
  port             = "${aws_lb_target_group.router-lb-tg1.port}"

  count = "${var.Counts["Infra"]}"
}


resource "aws_lb_target_group_attachment" "infra-node-tg2" {
  target_group_arn = "${aws_lb_target_group.router-lb-tg2.arn}"
  target_id        = "${aws_instance.infra-node.*.id[count.index]}"
  port             = "${aws_lb_target_group.router-lb-tg2.port}"

  count = "${var.Counts["Infra"]}"
}

