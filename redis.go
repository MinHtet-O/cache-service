package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	keyPrefix string
	rdb       *redis.Client
	ttl       time.Duration
}

var redisClient *RedisClient

func NewRedisClient(keyPrefix string) *RedisClient {

	if redisClient == nil {
		fmt.Println("client already created")
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB,
		})
		redisClient = &RedisClient{
			keyPrefix: "minhtet",
			rdb:       rdb,
			ttl:       time.Duration(300) * time.Second,
		}
	}
	return redisClient
}

func (c *RedisClient) Set(ctx context.Context, key string, value []byte) error {
	key = fmt.Sprintf("%s:%s", c.keyPrefix, key)
	err := c.rdb.WithContext(ctx).Set(key, value, c.ttl).Err()
	return err
}

func (c *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	key = fmt.Sprintf("%s:%s", c.keyPrefix, key)

	val, err := c.rdb.WithContext(ctx).Get(key).Result()

	return []byte(val), err
}

func main() {
	c := NewRedisClient("minhtet")

	ctx := context.Background()
	c.Set(ctx, "studentID", []byte("123"))
	val, err := c.Get(ctx, "studentName")

	if err == redis.Nil {
		fmt.Println("key does not exist")
	}
	fmt.Println(string(val))
}
