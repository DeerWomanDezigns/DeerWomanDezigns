resource "aws_lightsail_instance" "lightsail_instances" {
  for_each = var.lightsail_instances

  name              = each.value["name"]
  availability_zone = "us-east-2a"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"
  key_pair_name     = aws_lightsail_key_pair.lightsail_key_pairs[each.key].id
  user_data         = "${file("${path.module}/scripts/instance_setup.sh")}"
}

resource "aws_lightsail_static_ip" "lightsail_static_ips" {
  for_each = var.lightsail_instances

  name = "${each.value["name"]}_static"
}

resource "aws_lightsail_static_ip_attachment" "lightsail_static_ip_attachments" {
  for_each = var.lightsail_instances

  static_ip_name = aws_lightsail_static_ip.lightsail_static_ips[each.key].id
  instance_name  = aws_lightsail_instance.lightsail_instances[each.key].id
}

resource "aws_lightsail_instance_public_ports" "lightsail_instance_public_ports" {
  for_each = var.lightsail_instances

  instance_name = aws_lightsail_instance.lightsail_instances[each.key].id
  port_info {
    protocol  = "tcp"
    from_port = 22
    to_port   = 22
  }
  port_info {
    protocol  = "tcp"
    from_port = 80
    to_port   = 80
  }
  port_info {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443
  }
}

resource "aws_lightsail_key_pair" "lightsail_key_pairs" {
  for_each = var.lightsail_instances

  name       = "${each.key}_ssh"
  public_key = file("${path.module}/keys/${each.value["keyFile"]}")
}
