package flags

import (
	"blog_server/global"
	"blog_server/models"

	"github.com/sirupsen/logrus"
)

func MigrateDB() {
	err := global.DB.AutoMigrate(
		models.UserModel{},
		models.UserConfModel{},
		models.ArticleModel{},
		models.ArticleDiggModel{},
		models.CategoryModel{},
		models.CollectModel{},
		models.ImageModel{},
		models.UserArticleCollectModel{},
		models.UserArticleLookHistoryModel{},
		models.CommentModel{},
		models.BannerModel{},
		models.LogModel{},
		models.UserLoginModel{},
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败: %s", err)
	}
}
