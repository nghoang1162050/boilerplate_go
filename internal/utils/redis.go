package utils

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheClient struct {
	client *redis.Client
}

var RedisClient *CacheClient

func (c *CacheClient) NewRedisCache() error {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	url := fmt.Sprintf("redis://%s:%s/0?protocol=3", host, port)
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
	expiration, _ := strconv.Atoi(os.Getenv("REDIS_EXPIRATION"))
    return c.client.Set(ctx, key, value, time.Duration(expiration)*time.Minute).Err()
}

func (c *CacheClient) Delete(ctx context.Context, key string) error {
    return c.client.Del(ctx, key).Err()
}