variable "function_name" {
  description = "Name of the Lambda function"
  type        = string
  default     = "diet_app"
}

variable "function_role" {
  description = "Name of the Lambda function role"
  type        = string
  default     = "diet_app_role"
}

variable "ecr_repo_name" {
  description = "Name of the ECR repository"
  type        = string
  default     = "diet_app_ecr"
}

variable "api_gateway_name" {
  description = "Name of the API Gateway"
  type        = string
  default     = "diet_app_gw"
}

variable "ENV" {
  description = "ENV"
  type        = string
  default     = "PROD"
}

variable "AWS_LWA_PORT" {
  description = "Container Port"
  type        = number
  default     = 1323
}

variable "TURSO_URL" {
  description = "Turso URL"
  type        = string
  sensitive   = true
}

variable "TURSO_TOKEN" {
  description = "Turso Token"
  type        = string
  sensitive   = true
}

variable "NUTRITIONIX_APP_ID" {
  description = "Nutritionix App Id"
  type        = string
  sensitive   = true
}

variable "NUTRITIONIX_APP_KEY" {
  description = "Nutritionix App Key"
  type        = string
  sensitive   = true
}

variable "S3_ALLOWED_ORIGIN" {
  description = "S3 CORS Allowed Origin"
  type        = string
  sensitive   = true
}

variable "CLOUDFLARE_API_TOKEN" {
  description = "Cloudflare API Token"
  type        = string
  sensitive   = true
}

variable "CLOUDFLARE_ZONE_ID" {
  description = "Cloudflare DNS Zone ID"
  type        = string
  sensitive   = true
}
