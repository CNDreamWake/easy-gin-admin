package cache

import "github.com/redis/go-redis/v9"

var cache *redis.Client

func Initialization(conf *Cache) {
	cache = initRedis(&conf.Redis)
}

func GetCache() *redis.Client {
	return cache
}
