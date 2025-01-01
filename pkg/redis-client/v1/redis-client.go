package redisClient

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient defines the methods required for Redis interactions.
type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

var ctx = context.Background()

// redisClientInstance will hold the actual Redis client or a mock one
var redisClient RedisClient
var clientOnce sync.Once

// SetClient allows injection of a custom Redis client (for testing)
func SetClient(client RedisClient) {
	redisClient = client
}

func getRedisClient() RedisClient {
	clientOnce.Do(func() {
		if redisClient == nil {
			rdb := redis.NewClient(&redis.Options{
				Addr:     "redis:6379",
				Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
			})

			if rdb.ClientInfo(ctx).Err() != nil {
				panic(rdb.ClientInfo(ctx).Err())
			}
			redisClient = rdb
		}
	})

	return redisClient
}

func Set(key, value string) {
	client := getRedisClient()

	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	client := getRedisClient()

	value, err := client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return value
}
