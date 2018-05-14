resource "aws_efs_file_system" "persistence-storage" {
  count = "${var.EnableEfs}"

  creation_token = "${var.ProjectId}-openshift-storage"
  encrypted = true
  kms_key_id = "${aws_kms_key.efs-encryption-key.arn}"

  tags {
    Name = "${var.ProjectName} - Persistent Storage"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    Type = "persistence"
  }
}

resource "aws_efs_mount_target" "persistence-mount-targets" {
  count = "${(var.EnableEfs * aws_subnet.subnets-private.count)}"

  file_system_id = "${aws_efs_file_system.persistence-storage.id}"
  subnet_id      = "${aws_subnet.subnets-private.*.id[(count.index % aws_subnet.subnets-private.count)]}"
  security_groups = ["${aws_security_group.storage-sg.id}"]
}

resource "aws_kms_key" "efs-encryption-key" {
  count = "${var.EncryptEfs}"

  description             = "KMS encrytion key for OpenShift EFS encryption"
  deletion_window_in_days = 7

  tags {
    Name = "${var.ProjectName} - EFS Persistent Storage"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_kms_alias" "efs-encryption-key-alias" {
  count = "${var.EncryptEfs}"
  name          = "alias/${var.ProjectId}-efs-storage-key"
  target_key_id = "${aws_kms_key.efs-encryption-key.id}"
}

resource "aws_s3_bucket" "registry-storage" {
  count = "${var.RegistryS3}"

  bucket_prefix = "${var.ProjectId}-registry-"
  acl    = "private"
  region = "${var.Region}"
  force_destroy = true

  tags {
    Name = "${var.ProjectName} - Docker Registry Storage"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    Type = "registry"
  }
}