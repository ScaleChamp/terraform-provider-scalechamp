provider "scalechamp" {
    token = "9d1c7a6689bf9aeb7d5086067c1ce236"
    base_url = "http://api.scalablespace.net:3000"
}

resource "scalechamp_redis" "redis_cache" {
  name = "cache"
  plan = "hobby-100"
  cloud = "do"
  region = "fra1"
  whitelist = ["85.238.98.91"]
}

output "myredis" {
  value = "${scalechamp_redis.redis_cache.master_host} ${scalechamp_redis.redis_cache.password}"
}
