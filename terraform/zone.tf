data "aws_route53_zone" "existing-zone" {
  name = "${var.zone}"
}

data "aws_acm_certificate" "router-certificate" {
  domain   = "*.apps.${var.zone}"
  most_recent = true
}

resource "aws_route53_record" "router-record" {
  zone_id = "${data.aws_route53_zone.existing-zone.zone_id}"
  name    = "*.apps.${var.zone}"
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

resource "aws_route53_record" "bastion-record" {
  zone_id = "${data.aws_route53_zone.existing-zone.zone_id}"
  name    = "bastion.${data.aws_route53_zone.existing-zone.name}"
  type = "A"

  ttl = 300
  records = ["${aws_instance.bastion.public_ip}"]
}

resource "aws_route53_record" "internal-api-record" {
  zone_id = "${data.aws_route53_zone.existing-zone.zone_id}"
  name    = "internal-api.${data.aws_route53_zone.existing-zone.name}"
  type = "A"

  alias {
    name = "${aws_elb.internal-lb.dns_name}"
    evaluate_target_health = false
    zone_id = "${aws_elb.internal-lb.zone_id}"
  }
}
