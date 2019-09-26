provider "scalablespace" {
  apikey = "24870c2a-f8b2-4240-a41b-db98e2d76417"
}

resource "scalablespace_instance" "redis"{
  name = "myredis"
  plan = "Free"
  dc = "do-fra-1"
}

output "myredis" {
  value = "${scalablespace_instance.redis.port}"
}
