# Project name
variable "project" {
  type = "string"
  # Provide default for validation
  default = "Sprint1"
}

variable "project_id" {
  type = "string"
  default = "sprint1"
  description = "Project Id contains only lowercase alphanumerical characters."
}

# EC2 key to use
variable "key" {
  type = "string"
  default = "tobias@Codecentric"
  //default = "tobias@Home"
}

variable "node-types" {
  type = "map"
  default = {
    #"bastion" = "t2.micro"
    #"infrastructure" = "t2.large"
    #"application" = "t2.large"
    #"master" = "m4.xlarge"
    "bastion" = "t2.nano"
    "infrastructure" = "t2.medium"
    "application" = "t2.medium"
    "master" = "m4.large"
    # m5 = 10x faster network
  }
}

variable "zone" {
  type = "string"
  default = "cc-openshift.de"
}

variable "counts" {
  type = "map"
  default = {
    "master" = 2
    "infra" = 2
    "app" = 3
  }
}