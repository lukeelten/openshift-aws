module "application_node1" {
  source = "node"

  instance_type = "${var.node-types["application"]}"
  instance_ami = "${data.aws_ami.centos.id}"
  instance_key = "${var.key}"
  instance_name = "Application Node 1"
  instance_sg = ["${aws_security_group.nodes-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-private-1.id}"
  project = "${var.project}"
}

module "master_node1" {
  source = "node"

  instance_type = "${var.node-types["master"]}"
  instance_ami = "${data.aws_ami.centos.id}"
  instance_key = "${var.key}"
  instance_name = "Master Node 1"
  instance_sg = ["${aws_security_group.master-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-private-1.id}"
  project = "${var.project}"
}

module "infrastructure_node1" {
  source = "node"

  instance_type = "${var.node-types["infrastructure"]}"
  instance_ami = "${data.aws_ami.centos.id}"
  instance_key = "${var.key}"
  instance_name = "Infrastructure Node 1"
  instance_sg = ["${aws_security_group.infra-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-private-1.id}"
  project = "${var.project}"
}