package cacheRedis

import (
	"log"

	"github.com/go-redis/redis/v8"
)

func InitService() RedisServiceStruct {
	service := RedisServiceStruct{}
	service.client = initializeClient()
	return service
}

func initializeClient() *redis.Client {
	config, err := LoadConfig()
	if err != nil {
		log.Println(err)
		return nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetString(ConfigCacheAddress),
		Password: config.GetString(ConfigCachePassword),
		DB:       config.GetInt(ConfigCacheDb),
	})
	return rdb
}
