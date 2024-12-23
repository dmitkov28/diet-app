resource "aws_lambda_function" "lambda_func" {
  function_name = var.function_name
  package_type  = "Image"
  architectures = [local.architecture]
  image_uri     = "${data.aws_ecr_image.app_image.registry_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/${data.aws_ecr_image.app_image.repository_name}@${data.aws_ecr_image.app_image.image_digest}"
  role          = aws_iam_role.lambda_role.arn
  timeout       = 60
  memory_size   = 512

  environment {
    variables = {
      AWS_LWA_PORT = var.AWS_LWA_PORT
      TURSO_TOKEN  = var.TURSO_TOKEN
      TURSO_URL    = var.TURSO_URL
      NUTRITIONIX_APP_ID = var.NUTRITIONIX_APP_ID
      NUTRITIONIX_APP_KEY = var.NUTRITIONIX_APP_KEY
      ENV          = var.ENV
    }
  }

  depends_on = [null_resource.build_and_push, aws_ecr_repository.ecr_repo]
}

resource "aws_lambda_function_url" "lambda_function_url" {
  function_name      = aws_lambda_function.lambda_func.function_name
  authorization_type = "NONE"

}

resource "aws_iam_role" "lambda_role" {
  name = var.function_role

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

data "external" "uname" {
  program = ["sh", "-c", "echo '{\"arch\": \"'$(uname -m)'\"}'"]
}

locals {
  architecture = data.external.uname.result.arch
}
