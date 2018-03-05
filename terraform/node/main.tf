resource "aws_instance" "bastion" {
  ami           = "${var.instance_ami}"
  instance_type = "${var.instance_type}"
  key_name = "${var.instance_key}"

  vpc_security_group_ids = ["${var.instance_sg}"]
  subnet_id = "${var.instance_subnet}"

  tags {
    Name = "${var.project} - ${var.instance_name}"
    Project = "${var.project}"
  }
}