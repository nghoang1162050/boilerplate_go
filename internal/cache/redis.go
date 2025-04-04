package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CacheClient struct {
	client *redis.Client
}

var RedisClient *CacheClient

func (c *CacheClient) NewRedisCache() error {
	url := "redis://127.0.0.1:6379/0?protocol=3"
	options, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	c.client = redis.NewClient(options)

	return nil
}

func (c *CacheClient) Get(ctx context.Context, key string) (string, error) {
    return c.client.Get(ctx, key).Result()
}

func (c *CacheClient) Set(ctx context.Context, key, value string) error {
    return c.client.Set(ctx, key, value, 0).Err()
}

func (c *CacheClient) Delete(ctx context.Context, key string) error {
    return c.client.Del(ctx, key).Err()
}