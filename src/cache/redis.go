package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		WriteTimeout: 5 * time.Second,
		ReadTimeout: 5  * time.Second,
		DialTimeout: 5 * time.Second,
		Addr: "localhost:6379",
		Password: "password",
		DB: 0,
		PoolSize: 10,
		PoolTimeout: 500 * time.Millisecond,
	})
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}