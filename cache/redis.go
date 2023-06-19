package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"os"
	"time"
)

var client *redis.Client

func Init() {

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
}

// Get
// 如果key不存在，返回redis.Nil
func Get(key string) (string, error) {
	return client.Get(context.Background(), key).Result()
}

// Set
// redis: can't marshal struct {}
func Set(key string, value interface{}, expiration time.Duration) error {
	return client.Set(context.Background(), key, value, expiration).Err()
}

func Del(keys ...string) error {
	return client.Del(context.Background(), keys...).Err()
}
