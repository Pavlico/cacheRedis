package cacheRedis

import (
	"context"
	"time"
)

func SaveToCache(key string, value string) error {
	conf, err := LoadConfig()
	if err != nil {
		return err
	}
	expireTime := time.Duration(conf.GetInt32(ConfigCacheExpireTimeSeconds)) * time.Second
	rdb := InitService()
	ctx := context.Background()
	err = rdb.client.Set(ctx, key, value, expireTime).Err()
	return err
}
