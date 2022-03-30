package redisconfig

import "github.com/go-redis/redis/v8"

func InitializeRedisServer() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // default db
	})
	return rdb
}
