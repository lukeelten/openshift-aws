resource "aws_efs_file_system" "persistence-storage" {
  creation_token = "${var.project_id}-openshift-storage"
  encrypted = true
  kms_key_id = "${aws_kms_key.persistence-encryption-key.arn}"

  tags {
    Name = "${var.project} - Persistent Storage"
    Project = "${var.project}"
  }
}

resource "aws_kms_key" "persistence-encryption-key" {
  description             = "KMS encrytion key for OpenShift persistence encryption"
  deletion_window_in_days = 7

  tags {
    Name = "${var.project} - Persistent Storage"
    Project = "${var.project}"
  }
}

resource "aws_kms_alias" "persistence-encryption-key-alias" {
  name          = "alias/${var.project_id}-persistence-storage-key"
  target_key_id = "${aws_kms_key.persistence-encryption-key.id}"
}

resource "aws_efs_mount_target" "persistence-mt1" {
  file_system_id = "${aws_efs_file_system.persistence-storage.id}"
  subnet_id      = "${aws_subnet.subnet-private-1.id}"
  security_groups = ["${aws_security_group.storage-sg.id}"]
}

output "efs-id" {
  value = "${aws_efs_file_system.persistence-storage.id}"
}