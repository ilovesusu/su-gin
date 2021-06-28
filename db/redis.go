package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/ilovesusu/su-gin/configs"
)

var (
	SuRedis *redis.Client
)

func init() {
	SuRedis = redis.NewClient(&redis.Options{
		Addr:         configs.SuRedis.Host + ":" + configs.SuRedis.Port,
		Password:     configs.SuRedis.PassWord,
		DB:           configs.SuRedis.DatabasesName,
		PoolSize:     configs.SuRedis.POOL_SIZE,
		MinIdleConns: configs.SuRedis.MIN_IDLE_CONNS,
	})
}
