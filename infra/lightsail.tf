resource "aws_lightsail_instance" "lightsail_instance_backend" {
  name              = "dwd_backend"
  availability_zone = "us-east-2a"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"
  key_pair_name     = aws_lightsail_key_pair.lightsail_key_pair_backend.id
  user_data         = "${file("${path.module}/scripts/instance_setup.sh")}"
}

resource "aws_lightsail_instance" "lightsail_instance_frontend" {
  name              = "dwd_frontend"
  availability_zone = "us-east-2a"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"
  key_pair_name     = aws_lightsail_key_pair.lightsail_key_pair_frontend.id
  user_data         = "${file("${path.module}/scripts/instance_setup.sh")}"
}

resource "aws_lightsail_static_ip" "lightsail_static_ip_backend" {
  name = "backend_static"
}

resource "aws_lightsail_static_ip_attachment" "lightsail_static_ip_attachment_backend" {
  static_ip_name = aws_lightsail_static_ip.lightsail_static_ip_backend.id
  instance_name  = aws_lightsail_instance.lightsail_instance_backend.id
}

resource "aws_lightsail_static_ip" "lightsail_static_ip_frontend" {
  name = "frontend_static"
}

resource "aws_lightsail_static_ip_attachment" "lightsail_static_ip_attachment_frontend" {
  static_ip_name = aws_lightsail_static_ip.lightsail_static_ip_frontend.id
  instance_name  = aws_lightsail_instance.lightsail_instance_frontend.id
}

resource "aws_lightsail_instance_public_ports" "lightsail_instance_public_ports_backend" {
  instance_name = aws_lightsail_instance.lightsail_instance_backend.name

  port_info {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443
  }
}

resource "aws_lightsail_instance_public_ports" "lightsail_instance_public_ports_frontend" {
  instance_name = aws_lightsail_instance.lightsail_instance_frontend.name

  port_info {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443
  }
}

resource "aws_lightsail_key_pair" "lightsail_key_pair_backend" {
  name       = "backend_ssh"
  public_key = file("${path.module}/keys/backend.pub")
}

resource "aws_lightsail_key_pair" "lightsail_key_pair_frontend" {
  name       = "frontend_ssh"
  public_key = file("${path.module}/keys/frontend.pub")
}
