provider "aws" {
  region     = "us-east-1"
  access_key = "mock_access_key"
  secret_key = "mock_secret_key"
}

resource "aws_s3_bucket" "snykcon-2020-pictures" {
  bucket = "snykcon-2020-pictures"

  tags = {
    Name        = "Snykcon 2020 Pictures"
    Environment = "Production"
  }
}
