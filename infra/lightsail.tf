resource "aws_lightsail_instance" "lightsail_instance_backend" {
  name              = "dwd_backend"
  availability_zone = "us-east-2a"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"
  key_pair_name     = aws_lightsail_key_pair.lightsail_key_pair_backend.id
  user_data         = "${file("${path.module}/instance_setup.sh")}"
}

resource "aws_lightsail_instance" "lightsail_instance_frontend" {
  name              = "dwd_frontend"
  availability_zone = "us-east-2a"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"
  key_pair_name     = aws_lightsail_key_pair.lightsail_key_pair_frontend.id
  user_data         = "${file("${path.module}/instance_setup.sh")}"
}

resource "aws_lightsail_domain" "lightsail_domain_frontend" {
  domain_name = "deerwoman-dezigns.com"
}

resource "aws_lightsail_key_pair" "lightsail_key_pair_backend" {
  name       = "backend_ssh"
  public_key = file("${path.module}/keys/backend.pub")
}

resource "aws_lightsail_key_pair" "lightsail_key_pair_frontend" {
  name       = "frontend_ssh"
  public_key = file("${path.module}/keys/frontend.pub")
}
