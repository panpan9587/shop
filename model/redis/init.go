package redis

import (
	"demo/config"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{Addr: config.RedisDB})
}
