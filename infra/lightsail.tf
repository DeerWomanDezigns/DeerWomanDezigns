resource "aws_lightsail_instance" "lightsail_instance_backend" {
  name              = "dwd_backend"
  availability_zone = "us-east-2"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"
}

resource "aws_lightsail_instance" "lightsail_instance_frontend" {
  name              = "dwd_frontend"
  availability_zone = "us-east-2"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"
}
