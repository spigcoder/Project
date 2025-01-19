// models/category_model.go
package models

//用户分类表
type CategoryModel struct {
	Model
	Title     string    `gorm:"size:32" json:"title"`
	UserID    uint      `json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
}
