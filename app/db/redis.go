package db

import (
	"context"
	"fiber-starter/config"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetRedisAddr(),
		Password: config.GetRedisPassword(),
		DB:       0,
	})

	if err := RedisClient.Ping(RedisCtx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	log.Println("Connected to Redis")
}

func CloseRedis() {
	if err := RedisClient.Close(); err != nil {
		log.Fatal("Error closing redis: ", err)
	} else {
		log.Println("Redis connection closed")
	}
}

func DeleteCacheByID(id uint, cacheKeyTemplate string) error {
	cacheKey := fmt.Sprintf(cacheKeyTemplate, id)

	err := RedisClient.Del(RedisCtx, cacheKey).Err()
	if err != nil {
		log.Println("Error deleting cache: ", err)
		return err
	}

	log.Println("Cache deleted: ", cacheKey)
	return nil
}
