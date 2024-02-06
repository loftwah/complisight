resource "aws_s3_bucket" "non_compliant_bucket" {
  bucket = "non-compliant-bucket-${random_pet.name.id}"
}

resource "aws_s3_bucket_acl" "non_compliant_bucket_acl" {
  bucket = aws_s3_bucket.non_compliant_bucket.id
  acl    = "public-read"
}
