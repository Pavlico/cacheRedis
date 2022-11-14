package cacheRedis

import "github.com/go-redis/redis/v8"

type RedisServiceStruct struct {
	client *redis.Client
}
