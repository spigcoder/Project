package jwt

import (
	"blog_server/global"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// Token有一点就是不能主动设置失效时间，所以我们可以使用Redis黑名单机制来将其设置失效

type BlackList int8

const (
	UserLoseEfficacy BlackList = iota + 1
	AdminLoseEfficacy
	DevLoseEfficacy
	OtherLoseEfficacy
)

func (b BlackList) String() string {
	switch b {
	case UserLoseEfficacy:
		return "1"
	case AdminLoseEfficacy:
		return "2"
	case DevLoseEfficacy:
		return "3"
	}
	return "4"
}

func Parse(val string) BlackList {
	switch val {
	case "1":
		return UserLoseEfficacy
	case "2":
		return AdminLoseEfficacy
	case "3":
		return DevLoseEfficacy
	}
	return OtherLoseEfficacy
}

func LoseEfficacy(token string, value BlackList) (BlackList, error) {
	claim, err := ParseToken(token)
	key := fmt.Sprintf("black_list_%s", token)
	if err != nil {
		logrus.Error("ParseToken err: ", err)
	}
	expire := claim.ExpiresAt - time.Now().Unix()
	val, err := global.Redis.Set(context.Background(), key, value.String(), time.Duration(expire)*time.Second).Result()
	if err != nil {
		logrus.Error("Redis Set err: ", err)
	}
	return Parse(val), err
}

func IsBlackList(token string) (BlackList, bool) {
	key := fmt.Sprintf("black_list_%s", token)
	val, err := global.Redis.Get(context.Background(), key).Result()
	if err != nil {
		logrus.Error("Redis Get err: ", err)
		return OtherLoseEfficacy, false
	}
	if err == redis.Nil {
		return OtherLoseEfficacy, false
	}
	return Parse(val), true
}
