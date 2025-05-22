package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// type Cache interface {
// 	Set(ctx context.Context, key string, value any) error
// 	Get(ctx context.Context, key string) (string, error)
// 	Del(ctx context.Context, key string) error
// }

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter(client *redis.Client) *RedisAdapter {
	return &RedisAdapter{
		client: client,
	}
}

func (ra *RedisAdapter) Set(ctx context.Context, key string, value any) error {
	return ra.client.Set(ctx, key, value, 0).Err()
}
func (ra *RedisAdapter) Get(ctx context.Context, key string) (string, error) {
	return ra.client.Get(ctx, key).Result()
}
func (ra *RedisAdapter) Del(ctx context.Context, key string) error {
	return ra.client.Del(ctx, key).Err()
}
