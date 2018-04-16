# Project name
variable "project" {
  type = "string"
}

variable "project_id" {
  type = "string"
  description = "Project Id contains only lowercase alphanumerical characters."
}

variable "key" {
  description = "SSH key to use for EC2 instances"
  type = "string"
}

variable "node-types" {
  type = "map"
  #"bastion" = "t2.micro"
  #"infrastructure" = "t2.large"
  #"application" = "t2.large"
  #"master" = "m4.xlarge"
  # m5 = 10x faster network
}

variable "zone" {
  type = "string"
}

variable "counts" {
  type = "map"
}

variable "key_id" {
  type = "string"
  description = "Access key ID"
}

variable "secret_key" {
  type = "string"
  description = "Secret Access Key"
}

variable "region" {
  type = "string"
  description = "target AWS region"
}