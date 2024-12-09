variable "function_name" {
  description = "Name of the Lambda function"
  type        = string
  default     = "tf_debug"
}

variable "function_role" {
  description = "Name of the Lambda function role"
  type        = string
  default     = "tf_debug_role"
}

variable "ecr_repo_name" {
  description = "Name of the ECR repository"
  type        = string
  default     = "tf_debug_ecr"
}

variable "api_gateway_name" {
  description = "Name of the API Gateway"
  type        = string
  default     = "tf_debug_gw"
}

variable "AWS_LWA_PORT" {
  description = "Container Port"
  type        = number
  default     = 1323
}