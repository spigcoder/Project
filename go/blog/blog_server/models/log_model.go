// models/log_model.go
package models

import "blog_server/enum"

type LogModel struct {
	Model
	LogType     enum.LogType   `json:"logType"` // 日志类型 1 2 3
	Title       string         `gorm:"size:64" json:"title"`
	Content     string         `json:"content"`
	Level       enum.LogLevel  `json:"level"`                      // 日志级别 1 2 3
	UserID      uint           `json:"userID"`                     // 用户id
	UserModel   UserModel      `gorm:"foreignKey:UserID" json:"-"` // 用户信息
	IP          string         `gorm:"size:32" json:"ip"`
	Addr        string         `gorm:"size:64" json:"addr"`
	IsRead      bool           `json:"isRead"`              // 是否读取
	LoginStatus bool           `json:"loginStatus"`         // 是否登录
	UserName    string         `gorm:"size:32" json:"name"` // 用户名
	Pwd         string         `gorm:"size:32" json:"pwd"`  // 密码
	LoginType   enum.LoginType `json:"loginType"`           // 登录类型 1 2 3
}
