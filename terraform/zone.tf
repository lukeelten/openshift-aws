resource "aws_route53_zone" "openshift-zone" {
  name = "cc-openshift.de"

}

resource "aws_route53_record" "router-record" {
  zone_id = "${aws_route53_zone.openshift-zone.zone_id}"
  name    = "*.apps.cc-openshift.de"
  type    = "CNAME"
  ttl     = "60"

  records = [
    "${aws_lb.router-lb.dns_name}"
  ]
}

resource "aws_route53_record" "master-record" {
  zone_id = "${aws_route53_zone.openshift-zone.zone_id}"
  name    = "master.cc-openshift.de"
  type    = "CNAME"
  ttl     = "60"

  records = [
    "${aws_lb.master-lb.dns_name}"
  ]
}