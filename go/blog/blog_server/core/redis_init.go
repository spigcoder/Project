package core

import (
	"blog_server/global"
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	re := global.Config.Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     re.Addr,     // 不写默认就是这个
		Password: re.Password, // 密码,          // 密码
		DB:       re.DB,       // 默认是0
	})
	_, err := redisDB.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatal("redis连接失败")
	}

	logrus.Info("redis连接成功")
	return redisDB
}
