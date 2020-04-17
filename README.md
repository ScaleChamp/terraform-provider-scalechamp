# terraform-provider-scalechamp
[ScaleChamp](https://www.scalechamp.com) is a victorious managed databases provider

Example usage

```
provider "scalechamp" {
    token = "9d1c7a6689bf9aeb7d5086067c1ce236"
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
```