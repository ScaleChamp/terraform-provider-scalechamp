provider "scalablespace" {
}

resource "scalablespace_instance" "redis"{
  name = "myredis"
  plan = "Free"
  dc = "do-fra-1"
}

output "myredis" {
  value = "${scalablespace_instance.redis.port}"
}
