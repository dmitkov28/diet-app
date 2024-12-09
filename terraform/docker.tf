locals {
  source_files = flatten([
    fileset("../", "**/*.go"),
    fileset("../", "**/*.templ"),
    fileset("../", "go.{mod,sum}"),
    "Dockerfile.Lambda"
  ])
}

locals {
  source_hash = sha256(join("", [for f in local.source_files : filesha256("../${f}")]))
}

resource "null_resource" "build_and_push" {
  triggers = {
    source_hash = local.source_hash
  }

  provisioner "local-exec" {
    command = <<EOT
    docker build -t ${aws_ecr_repository.ecr_repo.repository_url}:latest -f ../Dockerfile.Lambda ../.
    aws ecr get-login-password --region ${data.aws_region.current.name} | docker login --username AWS --password-stdin ${aws_ecr_repository.ecr_repo.repository_url}
    docker push ${aws_ecr_repository.ecr_repo.repository_url}:latest
    EOT
  }

  depends_on = [aws_ecr_repository.ecr_repo]
}
