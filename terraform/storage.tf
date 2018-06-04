resource "aws_efs_file_system" "persistence-storage" {
  count = "${var.EnableEfs && !var.EncryptEfs ? 1 : 0}"
  creation_token = "${var.ProjectId}-openshift-storage"

  tags {
    Name = "${var.ProjectName} - Persistent Storage"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    Type = "persistence"
  }
}

resource "aws_efs_mount_target" "persistence-mount-targets" {
  count = "${aws_efs_file_system.persistence-storage.count * aws_subnet.subnets-private.count}"

  file_system_id = "${aws_efs_file_system.persistence-storage.id}"
  subnet_id      = "${aws_subnet.subnets-private.*.id[(count.index % aws_subnet.subnets-private.count)]}"
  security_groups = ["${aws_security_group.storage-sg.id}"]
}

resource "aws_efs_file_system" "persistence-storage-encrypted" {
  count = "${aws_kms_key.efs-encryption-key.count}"

  creation_token = "${var.ProjectId}-openshift-storage-encrypted"
  encrypted = true
  kms_key_id = "${aws_kms_key.efs-encryption-key.arn}"

  tags {
    Name = "${var.ProjectName} - Encrypted Persistent Storage"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
    Type = "persistence"
  }
}

resource "aws_efs_mount_target" "persistence-mount-targets-encrypted" {
  count = "${aws_efs_file_system.persistence-storage-encrypted.count * aws_subnet.subnets-private.count}"

  file_system_id = "${aws_efs_file_system.persistence-storage-encrypted.id}"
  subnet_id      = "${aws_subnet.subnets-private.*.id[(count.index % aws_subnet.subnets-private.count)]}"
  security_groups = ["${aws_security_group.storage-sg.id}"]
}

resource "aws_kms_key" "efs-encryption-key" {
  count = "${var.EncryptEfs && var.EnableEfs ? 1 : 0}"

  description             = "KMS encrytion key for OpenShift EFS encryption"
  deletion_window_in_days = 7

  tags {
    Name = "${var.ProjectName} - EFS Persistent Storage"
    Project = "${var.ProjectName}"
    ProjectId = "${var.ProjectId}"
  }
}

resource "aws_kms_alias" "efs-encryption-key-alias" {
  count = "${aws_kms_key.efs-encryption-key.count}"
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