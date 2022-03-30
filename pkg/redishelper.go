package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// every object is golang implements the empty interface i.e interface{}, i use that to reference a generic object in golang!
func RetrieveValueFromRedisClient(rdb *redis.Client, ctx *context.Context, key string) (string, bool, error) { // returns (object, isRetrieved, error)
	val, err := rdb.Get(*ctx, key).Result()

	var custom_error error = nil
	var is_retrieved bool = false

	if err == redis.Nil { // not an error
		fmt.Println("redis server: key does not exist")
	} else if err != nil { // is an err!
		fmt.Println("redis server: get failed", err)
		custom_error = errors.New(fmt.Sprintf("redis retrieve error: %v", err.Error()))
	} else if val == "" { // not an error
		fmt.Println("redis server: key fetched empty value")
	} else { // val != ""
		is_retrieved = true
	}

	return val, is_retrieved, custom_error
}

func InsertKeyValuePairInRedisClient(rdb *redis.Client, ctx *context.Context, key string, value interface{}) error {
	var err error = nil
	json, err := json.Marshal(value)
	// if err != nil {
	// 	panic(err)
	// }

	err = rdb.Set(*ctx, key, json, time.Hour*24).Err()
	if err != nil {
		// panic(err)
	} else {
		fmt.Println("redis server: data cached successfully")
	}
	return err
}
