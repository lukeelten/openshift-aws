resource "aws_iam_user_policy_attachment" "efs-user-policy" {
  user       = "${aws_iam_user.efs-user.name}"
  policy_arn = "${data.aws_iam_policy.efs-read-policy.arn}"
}

resource "aws_iam_user" "efs-user" {
  name = "${var.project_id}-efs-user"
  path = "/${var.project_id}/"
}

resource "aws_iam_access_key" "efs-user-key" {
  user = "${aws_iam_user.efs-user.name}"
}

output "efs-user-key" {
  value = "${aws_iam_access_key.efs-user-key.id}"
}

output "efs-user-secret" {
  value = "${aws_iam_access_key.efs-user-key.secret}"
}