output "compliant_s3_bucket_name" {
  description = "The name of the compliant S3 bucket."
  value       = aws_s3_bucket.compliant_bucket.bucket
}

output "non_compliant_s3_bucket_name" {
  description = "The name of the non-compliant S3 bucket."
  value       = aws_s3_bucket.non_compliant_bucket.bucket
}

output "rds_cluster_endpoint" {
  description = "The endpoint of the RDS cluster."
  value       = aws_rds_cluster.compliant_rds.endpoint
}

output "rds_cluster_master_username" {
  description = "The master username for the RDS cluster."
  value       = aws_rds_cluster.compliant_rds.master_username
}
