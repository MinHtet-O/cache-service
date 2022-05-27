package data

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var (
	redisClient *RedisClient
	addr        = "localhost:6379"
	passwd      = ""
)

type RedisClient struct {
	rdb       *redis.Client
	ttl       time.Duration
	keyPrefix string
}

type CacheNotFound struct{}

func (m *CacheNotFound) Error() string {
	return "cache not found"
}

func NewRedisClient(keyPrefix string) *RedisClient {
	if redisClient == nil {
		log.Println("redis client already created")
		rdb := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: passwd, // no password set
			DB:       0,      // use default DB,
		})
		redisClient = &RedisClient{
			rdb:       rdb,
			ttl:       time.Duration(300) * time.Second,
			keyPrefix: keyPrefix,
		}
	}
	return redisClient
}

func getKey(prefix string, key string) string {
	return fmt.Sprintf("%s:%s", prefix, key)
}
