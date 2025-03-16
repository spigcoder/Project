package global

import (
	"blog_server/conf"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
