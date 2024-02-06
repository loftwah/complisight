variable "aws_account_id" {
  description = "The AWS Account ID."
  type        = string
}

variable "master_username" {
  description = "Master username for the RDS cluster."
  type        = string
}

variable "master_password" {
  description = "Master password for the RDS cluster."
  type        = string
  sensitive   = true
}

variable "region" {
  description = "AWS region to deploy resources."
  type        = string
  default     = "ap-southeast-2"
}

variable "vpc_id" {
  description = "The VPC ID where resources will be deployed."
  type        = string
}

variable "subnet_ids" {
  description = "Subnet IDs for the RDS cluster."
  type        = list(string)
}
