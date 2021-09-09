package redisop

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

var cli *redis.Client

func TestMain(m *testing.M) {
	var err error
	op := &redis.Options{
		Addr:               "127.0.0.1:6379",
		DB:                 0,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: 500 * time.Millisecond,
		Password:           "654123",
	}
	cli, err = NewRedis(op)
	if err != nil {
		fmt.Printf("test main error=%v\n",err)
		return
	}
	m.Run()
}

func TestSet(t *testing.T) {
	err := cli.Set("key_test_set_001","value001",0).Err()
	if err != nil {
		t.Errorf("set error=%v\n",err)
	}
}

func TestGet(t *testing.T) {
	val,err := cli.Get("key_test_set_001").Result()
	if err != nil {
		t.Errorf("set error=%v\n",err)
	}
	t.Logf("value=%s\n",val)
}