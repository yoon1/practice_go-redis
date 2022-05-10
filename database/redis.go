package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var cache *redis.Client

func GetRedis() (*redis.Client, error) {
	if cache == nil {
		return nil, errors.New("redis is not connected")
	}
	return cache, nil
}

func ConnectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(context.TODO()).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	cache = client
}

func SetItem(ctx context.Context, r *redis.Client, key, val string) {
	err := r.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func GetItem(ctx context.Context, r *redis.Client, key string) string {
	val, err := r.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}

func GetAllKeys(ctx context.Context, r *redis.Client, key string) []string {
	keys := []string{}

	iter := r.Scan(ctx, 0, key, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}
