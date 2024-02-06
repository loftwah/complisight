resource "aws_s3_bucket" "non_compliant_bucket" {
  bucket_prefix = "non-compliant-bucket-"
}

resource "aws_s3_bucket_ownership_controls" "non_compliant_bucket_ownership_controls" {
  bucket = aws_s3_bucket.non_compliant_bucket.id

  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_public_access_block" "non_compliant_bucket_public_access_block" {
  bucket = aws_s3_bucket.non_compliant_bucket.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_acl" "non_compliant_bucket_acl" {
  depends_on = [
    aws_s3_bucket_ownership_controls.non_compliant_bucket_ownership_controls,
    aws_s3_bucket_public_access_block.non_compliant_bucket_public_access_block,
  ]

  bucket = aws_s3_bucket.non_compliant_bucket.id
  acl    = "public-read"
}
