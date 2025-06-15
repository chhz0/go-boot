package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

const redisURL = "redis://%s:%s@%s/%d"

type RedisOptions struct {
	User string
	Pass string
	Host string
	DB   int
}

func NewClient(opts RedisOptions) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(redisURL, opts.User, opts.Pass, opts.Host, opts.DB),
		Password: opts.Pass,
		DB:       opts.DB,
	})
}
