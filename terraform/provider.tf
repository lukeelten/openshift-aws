provider "aws" {
  region = "${var.Region}"
  access_key = "${var.KeyId}"
  secret_key = "${var.SecretKey}"

  version = "1.49"
}
