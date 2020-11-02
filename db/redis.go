package db

import (
	"github.com/go-redis/redis"
	"github.com/ilovesusu/su-gin/config"
)

var (
	SuRedis *redis.Client
)

func init() {
	SuRedis = redis.NewClient(&redis.Options{
		//todo 修改端口
		Addr:         config.SuRedis.Host + ":" + config.SuRedis.Port,
		Password:     config.SuRedis.PassWord,
		DB:           config.SuRedis.DatabasesName,
		PoolSize:     config.SuRedis.POOL_SIZE,
		MinIdleConns: config.SuRedis.MIN_IDLE_CONNS,
	})
}
