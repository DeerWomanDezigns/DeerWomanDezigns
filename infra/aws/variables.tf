variable "lightsail_instances" {
  type    = map(object({
    name = string
    keyFile = string
  }))
  default = {
    backend = {
      name = "dwd_backend"
      keyFile = "backend.pub"
    },
    frontend = {
      name = "dwd_frontend"
      keyFile = "frontend.pub"
    }
  }
}
