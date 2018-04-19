

resource "aws_instance" "bastion" {
  depends_on             = ["aws_internet_gateway.igw"]

  ami                    = "${data.aws_ami.centos.id}"
  instance_type          = "${var.Types["Bastion"]}"
  key_name               = "${var.SshKey}"

  subnet_id              = "${aws_subnet.subnet-public-1.id}"

  user_data              = "${file("scripts/bastion.sh")}"
  vpc_security_group_ids = ["${aws_security_group.bastion-sg.id}"]

  root_block_device {
    volume_type = "gp2"
    volume_size = 8
  }

  tags {
    Name = "${var.ProjectName} - Bastion"
    Project = "${var.ProjectName}"
    Type = "bastion"
    ProjectId = "${var.ProjectId}"
  }
}