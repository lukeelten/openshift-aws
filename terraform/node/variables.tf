variable "project" {
  type = "string"
  description = "Project Name"
}

variable "instance_name" {
  type = "string"
  description = "Instance Name"
}

variable "instance_type" {
  type = "string"
  description = "Instance Type"
}

variable "instance_ami" {
  type = "string"
  description = "AMI ID"
}

variable "instance_key" {
  type = "string"
  description = "EC2 key to use"
}

variable "instance_subnet" {
  type = "string"
  description = "Instance Subnet ID"
}

variable "instance_sg" {
  type = "list"
  description = "List of Security Group IDs"
}