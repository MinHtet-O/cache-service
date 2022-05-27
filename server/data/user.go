package data

import (
	"context"

	"github.com/go-redis/redis"
)

func (c *RedisClient) SetUser(ctx context.Context, key string, value []byte) error {
	err := c.rdb.WithContext(ctx).Set(key, value, c.ttl).Err()
	if err == redis.Nil {
		err = &CacheNotFound{}
	}
	return err
}

func (c *RedisClient) GetUser(ctx context.Context, key string) ([]byte, error) {
	val, err := c.rdb.WithContext(ctx).Get(key).Result()
	if err == redis.Nil {
		err = &CacheNotFound{}
	}
	return []byte(val), err
}
