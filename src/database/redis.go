package database

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		DialTimeout:  5 * time.Second,
		Addr:         "localhost:6379",
		Password:     "password",
		DB:           0,
		PoolSize:     10,
		PoolTimeout:  500 * time.Millisecond,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic("خطا در اتصال به Redis")
	}
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}

func Set[T any](key string, value T, duration time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return redisClient.Set(context.Background(), key, val, duration).Err()
}

func Get[T any](key string) (T, error) {
	var dest T = *new(T)
	val, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(val), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}
