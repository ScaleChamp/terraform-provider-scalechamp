provider "scalechamp" {
    token = "9d1c7a6689bf9aeb7d5086067c1ce236"
    base_url = "http://api.scalablespace.net:3000"
}

resource "scalechamp_redis" "redis_cache" {
  name = "cache"
  plan = "hobby-100"
  cloud = "do"
  region = "fra1"
}

resource "scalechamp_postgresql" "main_db" {
  name = "cache"
  plan = "hobby-100"
  cloud = "do"
  region = "fra1"
}

output "main_db1" {
  value = "${scalechamp_postgresql.main_db.master_host} ${scalechamp_postgresql.main_db.password}"
}


output "myredis" {
  value = "${scalechamp_redis.redis_cache.master_host} ${scalechamp_redis.redis_cache.password}"
}
