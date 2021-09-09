package redisop

import (
	"github.com/go-redis/redis"
)

// NewRedis /**
func NewRedis(conf *redis.Options) (*redis.Client, error) {
	cli := redis.NewClient(conf)
	_, err := cli.Ping().Result()
	if err != nil {
		return nil, err
	}
	return cli, nil
}
