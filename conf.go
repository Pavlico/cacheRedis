package cacheRedis

import "github.com/spf13/viper"

const configFileType = "json"
const configPath = "."
const configFileName = "cache"
const ConfigCacheAddress = "Addr"
const ConfigCachePassword = "Password"
const ConfigCacheDb = "DB"
const ConfigCacheExpireTimeSeconds = "ExpireTimeSeconds"

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configFileName)
	v.SetConfigType(configFileType)
	err := v.ReadInConfig()
	if err != nil {
		return v, err
	}
	return v, nil
}
