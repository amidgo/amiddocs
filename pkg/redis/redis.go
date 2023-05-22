package redis

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*redis.Client
}

func New(addr string, opts ...Option) *Redis {
	rd := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	redis := &Redis{rd}
	for _, opt := range opts {
		opt(redis)
	}
	return redis
}

type Option func(rd *Redis)
