package cache

import (
	"fmt"
	"tasoryx/config"
	"time"

	"github.com/redis/go-redis/v9"
)

func Init(opt config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("$s:$s", opt.Host, opt.Port),
		Password:     opt.Password,
		DB:           opt.DB,
		PoolSize:     opt.PoolSize,
		DialTimeout:  time.Second * time.Duration(opt.DialTimeout),
		ReadTimeout:  time.Second * time.Duration(opt.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(opt.WriteTimeout),
	})
	return client
}

// TODO: need to review
