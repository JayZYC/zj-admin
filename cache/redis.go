package cache

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
)

const (
	Token = "token_" // 自定义登录过期时间
	Perm  = "perm_"  // 用户按钮权限
)

var client *redis.Client

func Init() {

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_address"),
		Password: os.Getenv("redis_password"),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot connect redis")
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

// Expire
// 刷新key过期时间
func Expire(key string, expiration time.Duration) {
	client.Expire(context.Background(), key, expiration)
}

// Get
// 如果key不存在，返回redis.Nil
func GetArr(key string) (interface{}, error) {
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	var value interface{}
	err = json.Unmarshal([]byte(val), &value)
	return value, err
}

// SetArr
// 序列化数组或结构体后存入
func SetArr(key string, value interface{}, expiration time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(context.Background(), key, val, expiration).Err()
}
