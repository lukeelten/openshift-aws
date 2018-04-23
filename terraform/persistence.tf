resource "aws_efs_file_system" "persistence-storage" {
  count = "${var.EnableEfs}"

  creation_token = "${var.ProjectId}-openshift-storage"
  encrypted = true
  kms_key_id = "${aws_kms_key.persistence-encryption-key.arn}"

  tags {
    Name = "${var.ProjectName} - Persistent Storage"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    Type = "persistence"
  }
}

resource "aws_efs_mount_target" "persistence-mt1" {
  count = "${var.EnableEfs}"

  file_system_id = "${aws_efs_file_system.persistence-storage.id}"
  subnet_id      = "${aws_subnet.subnet-private-1.id}"
  security_groups = ["${aws_security_group.storage-sg.id}"]
}

resource "aws_kms_key" "persistence-encryption-key" {
  count = "${var.EncryptEfs}"

  description             = "KMS encrytion key for OpenShift persistence encryption"
  deletion_window_in_days = 7

  tags {
    Name = "${var.ProjectName} - Persistent Storage"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_kms_alias" "persistence-encryption-key-alias" {
  count = "${var.EncryptEfs}"
  name          = "alias/${var.ProjectId}-persistence-storage-key"
  target_key_id = "${aws_kms_key.persistence-encryption-key.id}"
}