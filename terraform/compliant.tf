provider "aws" {
  region = var.region
}

resource "random_pet" "name" {}

resource "aws_s3_bucket" "compliant_bucket" {
  bucket_prefix = "compliant-bucket-${random_pet.name.id}"
}

resource "aws_s3_bucket_acl" "compliant_bucket_acl" {
  bucket = aws_s3_bucket.compliant_bucket.id
  acl    = "private"
}

resource "aws_s3_bucket_server_side_encryption_configuration" "compliant_bucket_sse" {
  bucket = aws_s3_bucket.compliant_bucket.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_kms_key" "rds_encryption_key" {
  description = "KMS key for RDS encryption"
}

resource "aws_rds_cluster" "compliant_rds" {
  engine                     = "aurora-postgresql"
  engine_version             = "13.4"
  cluster_identifier         = "compliant-rds-cluster-${random_pet.name.id}" # Use cluster_identifier
  master_username            = var.master_username
  master_password            = var.master_password
  skip_final_snapshot        = true
  storage_encrypted          = true
  kms_key_id                 = aws_kms_key.rds_encryption_key.arn
  db_subnet_group_name       = aws_db_subnet_group.example.name
  vpc_security_group_ids     = [aws_security_group.example.id]
}

resource "aws_db_subnet_group" "example" {
  name       = "my-db-subnet-group-${random_pet.name.id}"
  subnet_ids = var.subnet_ids
}

resource "aws_security_group" "example" {
  name        = "my-rds-sg-${random_pet.name.id}"
  description = "Managed by Terraform"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
