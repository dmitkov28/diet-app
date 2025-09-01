resource "aws_cloudfront_origin_access_control" "s3_oac" {
  name                              = "dietapp-s3-oac"
  description                       = "OAC for S3 static assets"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_cloudfront_cache_policy" "caching_disabled" {
  name        = "DisableCaching"
  comment     = "Policy with caching disabled"
  min_ttl     = 0
  max_ttl     = 0
  default_ttl = 0

  parameters_in_cache_key_and_forwarded_to_origin {
    enable_accept_encoding_gzip   = false
    enable_accept_encoding_brotli = false

    cookies_config {
      cookie_behavior = "none"
    }

    headers_config {
      header_behavior = "none"
    }

    query_strings_config {
      query_string_behavior = "none"
    }
  }
}

resource "aws_cloudfront_cache_policy" "static_assets_caching" {
  name        = "StaticAssetsCaching"
  comment     = "Policy with caching enabled for static assets like CSS, JS, images"
  min_ttl     = 0
  max_ttl     = 31536000
  default_ttl = 86400

  parameters_in_cache_key_and_forwarded_to_origin {
    enable_accept_encoding_gzip   = true
    enable_accept_encoding_brotli = true

    cookies_config {
      cookie_behavior = "none"
    }

    headers_config {
      header_behavior = "none"
    }

    query_strings_config {
      query_string_behavior = "none"
    }
  }
}

resource "aws_cloudfront_origin_request_policy" "api_gateway" {
  name    = "AllViewerExceptHostHeader"
  comment = "Policy to forward all parameters in viewer requests except for the Host header"

  cookies_config {
    cookie_behavior = "all"
  }

  headers_config {
    header_behavior = "allExcept"
    headers {
      items = ["Host"]
    }
  }

  query_strings_config {
    query_string_behavior = "all"
  }
}

resource "aws_cloudfront_response_headers_policy" "api_gateway" {
  name    = "SimpleCORS"
  comment = "Allows all origins for simple CORS requests"

  cors_config {
    access_control_allow_credentials = false

    access_control_allow_headers {
      items = ["*"]
    }

    access_control_allow_methods {
      items = ["GET", "HEAD", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"]
    }

    access_control_allow_origins {
      items = ["*"]
    }

    origin_override = true
  }
}

resource "aws_cloudfront_distribution" "cloudfront_dist" {
  aliases = [
    "diet.dimitarmitkov.com",
  ]
  enabled         = true
  http_version    = "http2"
  is_ipv6_enabled = true
  tags            = {}
  tags_all        = {}

  wait_for_deployment = true
  web_acl_id          = null

  default_cache_behavior {
    allowed_methods = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods  = ["GET", "HEAD"]

    target_origin_id       = "dietapp"
    compress               = true
    viewer_protocol_policy = "redirect-to-https"

    cache_policy_id            = aws_cloudfront_cache_policy.caching_disabled.id
    origin_request_policy_id   = aws_cloudfront_origin_request_policy.api_gateway.id
    response_headers_policy_id = aws_cloudfront_response_headers_policy.api_gateway.id
  }

  ordered_cache_behavior {
    path_pattern     = "/static/*"
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "s3-static"

    compress               = true
    viewer_protocol_policy = "redirect-to-https"

    cache_policy_id = aws_cloudfront_cache_policy.static_assets_caching.id
  }

  origin {
    connection_attempts      = 3
    connection_timeout       = 10
    domain_name              = "${aws_api_gateway_rest_api.api_gateway.id}.execute-api.${data.aws_region.current.name}.amazonaws.com"
    origin_access_control_id = null
    origin_id                = "dietapp"
    origin_path              = "/prod"

    custom_origin_config {
      http_port                = 80
      https_port               = 443
      origin_keepalive_timeout = 5
      origin_protocol_policy   = "https-only"
      origin_read_timeout      = 30
      origin_ssl_protocols     = ["TLSv1.2"]
    }
  }


  origin {
    domain_name              = aws_s3_bucket.s3_bucket.bucket_regional_domain_name
    origin_id                = "s3-static"
    origin_access_control_id = aws_cloudfront_origin_access_control.s3_oac.id


  }

  restrictions {
    geo_restriction {
      locations        = []
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn            = aws_acm_certificate.cert.arn
    cloudfront_default_certificate = false
    iam_certificate_id             = null
    minimum_protocol_version       = "TLSv1.2_2021"
    ssl_support_method             = "sni-only"
  }
}
