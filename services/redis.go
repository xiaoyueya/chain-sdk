package services

import (
	"github.com/go-redis/redis"
	"github.com/xiaoyueya/chain-sdk/common"
	"github.com/xiaoyueya/chain-sdk/redisop"
	"time"
)

var redisService *RedisService

type RedisService struct {
	client *redis.Client
}

func GetRawService() *RedisService {
	return redisService
}

func InitRedisService(redisCfg *common.RedisConfig) error {
	var err error
	redisService, err = NewRedisService(redisCfg)
	if err != nil {
		return err
	}
	return nil
}

func NewRedisService(redisCfg *common.RedisConfig) (*RedisService, error) {
	client, er := redisop.NewRedis(&redis.Options{
		Addr:               redisCfg.Addr,
		DB:                 redisCfg.DB,
		DialTimeout:        time.Duration(redisCfg.DialTimeout) * time.Second,
		ReadTimeout:        time.Duration(redisCfg.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(redisCfg.WriteTimeout) * time.Second,
		PoolSize:           redisCfg.PoolSize,
		PoolTimeout:        time.Duration(redisCfg.PoolTimeout) * time.Second,
		IdleTimeout:        time.Duration(redisCfg.IdleTimeout) * time.Millisecond,
		IdleCheckFrequency: time.Duration(redisCfg.IdleCheckFrequency) * time.Millisecond,
		Password:           redisCfg.Password,
	},
	)

	if er != nil {
		return nil, er
	}
	return &RedisService{client: client}, nil
}

func (service *RedisService) GetClient() *redis.Client {
	return service.client
}
