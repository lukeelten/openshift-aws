
resource "aws_route53_record" "router-record" {
  zone_id = "${data.aws_route53_zone.existing-zone.zone_id}"
  name    = "*.apps.${var.Zone}"
  type = "A"

  alias {
    name = "${aws_lb.router-lb.dns_name}"
    evaluate_target_health = false
    zone_id = "${aws_lb.router-lb.zone_id}"
  }
}

resource "aws_route53_record" "master-record" {
  zone_id = "${data.aws_route53_zone.existing-zone.zone_id}"
  name    = "master.${data.aws_route53_zone.existing-zone.name}"
  type = "A"

  alias {
    name = "${aws_lb.master-lb.dns_name}"
    evaluate_target_health = false
    zone_id = "${aws_lb.master-lb.zone_id}"
  }
}

resource "aws_route53_record" "internal-api-record" {
  zone_id = "${data.aws_route53_zone.existing-zone.zone_id}"
  name    = "internal-api.${data.aws_route53_zone.existing-zone.name}"
  type = "A"

  count = "${aws_lb.internal-lb.count}"

  alias {
    name = "${aws_lb.internal-lb.dns_name}"
    evaluate_target_health = false
    zone_id = "${aws_lb.internal-lb.zone_id}"
  }
}

resource "aws_route53_record" "internal-api-record-single-master" {
  zone_id = "${data.aws_route53_zone.existing-zone.zone_id}"
  name    = "internal-api.${data.aws_route53_zone.existing-zone.name}"
  type = "A"

  ttl = 300

  count = "${aws_lb.internal-lb.count > 0 ? 0 : 1}"
  records = ["${aws_instance.master-node.*.private_ip}"]
}