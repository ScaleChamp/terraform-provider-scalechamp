# terraform-provider-scalechamp
[ScaleChamp](https://www.scalechamp.com) is a victorious managed databases provider

Example usage

```
provider "scalechamp" {
    token = "<token>"
}

resource "scalechamp_postgresql" "main_db" {
  name = "main_db"
  plan = "hobby-100"
  cloud = "do"
  region = "fra1"
  whitelist = ["<ip|subnet>"]
}

resource "scalechamp_redis" "redis_cache" {
  name = "cache"
  plan = "hobby-100"
  cloud = "do"
  region = "fra1"
  whitelist = ["<ip|subnet>"]
}

output "mypg" {
  value = "${scalechamp_postgresql.main_db.master_host} ${scalechamp_postgresql.main_db.password}"
}

output "myredis" {
  value = "${scalechamp_redis.redis_cache.master_host} ${scalechamp_redis.redis_cache.password}"
}
```
