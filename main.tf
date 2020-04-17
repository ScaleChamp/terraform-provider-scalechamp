provider "scalechamp" {
    token = ""
}

resource "scalechamp_redis" "redis_cache" {
  name = "cache"
  plan = "hobby-100"
  cloud = "do"
  region = "fra-1"
  whitelist = ["85.238.98.98"]
}

output "myredis" {
  value = "${scalechamp_redis.redis_cache.master_host} ${scalechamp_redis.redis_cache.password}"
}
