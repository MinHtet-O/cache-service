package data

import (
	"context"

	"github.com/go-redis/redis"
)

func (c *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	key = getKey(c.keyPrefix, key)
	err := c.rdb.WithContext(ctx).Set(key, value, c.ttl).Err()
	if err == redis.Nil {
		err = &CacheNotFound{}
	}
	return err
}

func (c *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	key = getKey(c.keyPrefix, key)
	val, err := c.rdb.WithContext(ctx).Get(key).Result()
	if err == redis.Nil {
		err = &CacheNotFound{}
	}
	return []byte(val), err
}
