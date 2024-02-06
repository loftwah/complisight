terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.0" # Make sure this is compatible with your resources
    }
  }
  required_version = ">= 0.12" # Ensure your Terraform version is also compatible
}
