resource "aws_instance" "node" {
  ami           = "${var.instance_ami}"
  instance_type = "${var.instance_type}"
  key_name = "${var.instance_key}"

  vpc_security_group_ids = ["${var.instance_sg}"]
  subnet_id = "${var.instance_subnet}"

  root_block_device {
    volume_type = "standard"
    volume_size = "${var.root_size}"
  }

  tags {
    Name = "${var.project} - ${var.instance_name}"
    Project = "${var.project}"
  }
}

output "node-id" {
  value = "${aws_instance.node.id}"
}

output "internal-dns" {
  value = "${aws_instance.node.private_dns}"
}

output "internal-ip" {
  value = "${aws_instance.node.private_ip}"
}

output "external-dns" {
  value = "${aws_instance.node.public_dns}"
}

output "external-ip" {
  value = "${aws_instance.node.public_ip}"
}