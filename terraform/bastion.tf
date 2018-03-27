

resource "aws_instance" "bastion" {
  depends_on             = ["aws_internet_gateway.igw"]

  ami                    = "${data.aws_ami.centos.id}"
  instance_type          = "${var.node-types["bastion"]}"
  key_name               = "${var.key}"

  subnet_id              = "${aws_subnet.subnet-public-1.id}"

  user_data              = "${file("scripts/bastion.sh")}"
  vpc_security_group_ids = ["${aws_security_group.bastion-sg.id}"]

  root_block_device {
    volume_type = "gp2"
    volume_size = 8
  }

  tags {
    Name = "${var.project} - Bastion"
    Project = "${var.project}"
    Type = "bastion"
  }
}

output "bastion-dns" {
  value = "${aws_instance.bastion.public_dns}"
}

output "bastion-ip" {
  value = "${aws_instance.bastion.public_ip}"
}