package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func SetValue(redisClient *redis.Client, key string, value string) error {
	err := redisClient.Set(key, value, 1*time.Hour).Err()
	if err != nil {
		fmt.Println("redis set value error -- ", err)
	}
	return err
}

func GetValue(redisClient *redis.Client, key string) (string, error) {
	val, err := redisClient.Get(key).Result()
	if err == nil {
		return val, err
	}
	return "", err
}
