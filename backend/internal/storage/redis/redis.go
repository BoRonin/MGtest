package redis

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(address string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: address,
		DB:   0,
	})
}
