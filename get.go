package github.com/Pavlico/rediscache/cacheRedis

import (
	"context"
)

func GetByKey(key string) (string, error) {
	rdb := InitService()
	ctx := context.Background()
	val, err := rdb.client.Get(ctx, key).Result()
	return val, err
}
