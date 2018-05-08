resource "aws_iam_role" "master-role" {
  name               = "${var.ProjectId}-master-role"
  assume_role_policy = "${file("assets/role.json")}"
}

resource "aws_iam_role" "node-role" {
  name               = "${var.ProjectId}-node-role"
  assume_role_policy = "${file("assets/role.json")}"
}

resource "aws_iam_policy" "master-ec2-policy" {
  name        = "${var.ProjectId}-master-ec2-policy"
  description = "A test policy"
  policy      = "${file("assets/master-ec2.json")}"
}

resource "aws_iam_policy" "master-elb-policy" {
  name        = "${var.ProjectId}-master-elb-policy"
  description = "A test policy"
  policy      = "${file("assets/master-elb.json")}"
}

resource "aws_iam_policy" "nodes-policy" {
  name        = "${var.ProjectId}-node-policy"
  description = "A test policy"
  policy      = "${file("assets/nodes.json")}"
}

resource "aws_iam_policy_attachment" "master-ec2-attach" {
  name       = "${var.ProjectId}-master-ec2-attach"
  roles      = ["${aws_iam_role.master-role.name}"]
  policy_arn = "${aws_iam_policy.master-ec2-policy.arn}"
}

resource "aws_iam_policy_attachment" "master-elb-attach" {
  name       = "${var.ProjectId}-master-elb-attach"
  roles      = ["${aws_iam_role.master-role.name}"]
  policy_arn = "${aws_iam_policy.master-elb-policy.arn}"
}

resource "aws_iam_policy_attachment" "nodes-attach" {
  name       = "${var.ProjectId}-nodes-attach"
  roles      = ["${aws_iam_role.node-role.name}"]
  policy_arn = "${aws_iam_policy.nodes-policy.arn}"
}

resource "aws_iam_instance_profile" "node-profile" {
  name  = "${var.ProjectId}-node-profile"
  role = "${aws_iam_role.node-role.name}"
}

resource "aws_iam_instance_profile" "master-profile" {
  name  = "${var.ProjectId}-master-profile"
  role = "${aws_iam_role.master-role.name}"
}