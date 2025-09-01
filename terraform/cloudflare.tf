resource "cloudflare_dns_record" "diet_app_cname" {
  zone_id = var.CLOUDFLARE_ZONE_ID
  name    = "diet"
  type    = "CNAME"
  content = aws_cloudfront_distribution.cloudfront_dist.domain_name
  ttl     = 1
  proxied = false

  depends_on = [aws_cloudfront_distribution.cloudfront_dist]
}
