package redis

import (
	"github.com/bighuangbee/gomod/config"
	"github.com/bighuangbee/gomod/loger"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func init(){
	setup()
}

func setup(){

	Redis = redis.NewClient(&redis.Options{
		Addr:     config.ConfigData.RedisAddr,
		Password: config.ConfigData.RedisPassword,
		DB:       config.ConfigData.RedisDefaultDB,
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		panic("Redis Client SetUp Failed!" + err.Error())
		return
	}

	loger.Info("Redis Client SetUp Success...")
}