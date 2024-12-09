resource "aws_ecr_repository" "ecr_repo" {
  name         = var.ecr_repo_name
  force_delete = true
}

data "aws_ecr_image" "python_image" {
  repository_name = aws_ecr_repository.ecr_repo.name
  image_tag       = "latest"

  depends_on = [null_resource.build_and_push]
}