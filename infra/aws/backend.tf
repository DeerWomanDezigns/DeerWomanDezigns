terraform {
  backend "s3" {
    bucket = "deerwoman-terraform-state"
    key    = "terraform.tfstate"
    region = "us-east-2"
  }
}
