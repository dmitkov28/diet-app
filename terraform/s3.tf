resource "aws_s3_bucket" "s3_bucket" {
  bucket        = "dietapp-static"
  force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "block_settings" {
  bucket = aws_s3_bucket.s3_bucket.id

  block_public_acls       = true
  ignore_public_acls      = true
  block_public_policy     = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_cors_configuration" "static_assets" {
  bucket = aws_s3_bucket.s3_bucket.id

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "HEAD"]
    allowed_origins = [var.S3_ALLOWED_ORIGIN]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }
}

resource "aws_s3_bucket_policy" "static_assets" {
  bucket     = aws_s3_bucket.s3_bucket.id
  depends_on = [aws_s3_bucket_public_access_block.block_settings]
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowCloudFrontServicePrincipal"
        Effect = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action   = "s3:GetObject"
        Resource = "${aws_s3_bucket.s3_bucket.arn}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = aws_cloudfront_distribution.cloudfront_dist.arn
          }
        }
      }
    ]
  })
}
