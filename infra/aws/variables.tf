variable "lightsail_instances" {
  type    = map(object({
    name = string
    keyFile = string
  }))
  default = {
    backend = {
      name = "dwd_backend"
      keyFile = "backend.pub"
      lightsailBundleId = "micro_2_0"
    },
    frontend = {
      name = "dwd_frontend"
      keyFile = "frontend.pub"
      lightsailBundleId = "nano_2_0"
    }
  }
}
