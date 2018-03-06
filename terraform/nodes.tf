module "application_node1" {
  source = "node"

  instance_type = "${var.node-types["application"]}"
  instance_ami = "${data.aws_ami.fedora_atomic.id}"
  instance_key = "${var.key}"
  instance_name = "Application Node 1"
  instance_sg = ["${aws_security_group.nodes-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-public-1.id}"
  project = "${var.project}"
}

module "application_node2" {
  source = "node"

  instance_type = "${var.node-types["application"]}"
  instance_ami = "${data.aws_ami.fedora_atomic.id}"
  instance_key = "${var.key}"
  instance_name = "Application Node 2"
  instance_sg = ["${aws_security_group.nodes-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-public-2.id}"
  project = "${var.project}"
}

module "application_node3" {
  source = "node"

  instance_type = "${var.node-types["application"]}"
  instance_ami = "${data.aws_ami.fedora_atomic.id}"
  instance_key = "${var.key}"
  instance_name = "Application Node 3"
  instance_sg = ["${aws_security_group.nodes-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-public-3.id}"
  project = "${var.project}"
}

module "master_node1" {
  source = "node"

  instance_type = "${var.node-types["master"]}"
  instance_ami = "${data.aws_ami.fedora_atomic.id}"
  instance_key = "${var.key}"
  instance_name = "Master Node 1"
  instance_sg = ["${aws_security_group.master-sg.id}", "${aws_security_group.etcd-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-public-3.id}"
  project = "${var.project}"
}

module "master_node2" {
  source = "node"

  instance_type = "${var.node-types["master"]}"
  instance_ami = "${data.aws_ami.fedora_atomic.id}"
  instance_key = "${var.key}"
  instance_name = "Master Node 2"
  instance_sg = ["${aws_security_group.master-sg.id}", "${aws_security_group.etcd-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-public-2.id}"
  project = "${var.project}"
}

module "infrastructure_node1" {
  source = "node"

  instance_type = "${var.node-types["infrastructure"]}"
  instance_ami = "${data.aws_ami.fedora_atomic.id}"
  instance_key = "${var.key}"
  instance_name = "Infrastructure Node 1"
  instance_sg = ["${aws_security_group.infra-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-public-1.id}"
  project = "${var.project}"
}

module "infrastructure_node2" {
  source = "node"

  instance_type = "${var.node-types["infrastructure"]}"
  instance_ami = "${data.aws_ami.fedora_atomic.id}"
  instance_key = "${var.key}"
  instance_name = "Infrastructure Node 2"
  instance_sg = ["${aws_security_group.infra-sg.id}", "${aws_security_group.allow-all-sg.id}"]
  instance_subnet = "${aws_subnet.subnet-public-2.id}"
  project = "${var.project}"
}