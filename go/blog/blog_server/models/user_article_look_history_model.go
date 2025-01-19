// models/user_article_look_history_model.go
package models

//用户访问历史表
type UserArticleLookHistoryModel struct {
  Model
  UserID       uint         `json:"userID"`
  UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
  ArticleID    uint         `json:"articleID"`
  ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`
}
