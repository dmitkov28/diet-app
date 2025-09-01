resource "aws_acm_certificate" "cert" {
  provider    = aws.us-east-1
  domain_name = "diet.dimitarmitkov.com"

  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}
