package db

import (
	"context"
	"fiber-scaffold/common/logging"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var Rdb *redis.Client

func SetupRedis() {
	Rdb = redis.NewClient(&redis.Options{
		//相关信息放到config 本地调试
		Addr:     viper.GetString("REDIS.ADDR"),
		Password: viper.GetString("REDIS.PASS"),
		DB:       viper.GetInt("REDIS.DB"),
	})
	ping := Rdb.Ping(context.Background()).Val()
	if ping != "PONG" {
		logging.Fatal("连接失败")
	} else {
		logging.Info("connect redis success")
	}
}

var ctx = context.Background()

func Get(key string) string {
	val, err := Rdb.Get(ctx, key).Result()
	//fmt.Println("val", val, err)
	if err != nil {
		logging.Info(key, "数据为空", err)
	}
	return val
}

func Set(key, data string, exp time.Duration) {
	err := Rdb.Set(ctx, key, data, exp).Err()
	if err != nil {
		logging.Info("set数据出错", err)
	}
}
