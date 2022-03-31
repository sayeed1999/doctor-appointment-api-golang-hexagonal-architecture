package redisconfig

import "github.com/go-redis/redis/v8"

func InitializeRedisServer(addr string, pass string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	return rdb
}
