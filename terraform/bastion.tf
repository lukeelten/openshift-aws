resource "aws_instance" "bastion" {
  #ami           = "${data.aws_ami.centos.id}"
  ami = "ami-3d1e7352"
  instance_type = "${var.node-types["bastion"]}"
  key_name = "${var.key}"

  vpc_security_group_ids = ["${aws_security_group.bastion-sg.id}"]
  subnet_id = "${aws_subnet.subnet-public-1.id}"

  tags {
    Name = "${var.project} - Bastion"
    Project = "${var.project}"
  }
}

output "bastion-dns" {
  value = "${aws_instance.bastion.public_dns}"
}

output "bastion-ip" {
  value = "${aws_instance.bastion.public_ip}"
}