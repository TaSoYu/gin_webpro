package config

import (
	"context"
	"exchangeapp/global"
	"github.com/go-redis/redis/v8"
	"log"
)

func initRedis() {
	addr := AppConfig.Redis.Addr
	db := AppConfig.Redis.DB
	password := AppConfig.Redis.Password

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})

	_, err := RedisClient.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal(err)
	}

	global.RedisDB = RedisClient
}
