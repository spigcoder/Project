package logService

import (
	"blog_server/core"
	"blog_server/enum"
	"blog_server/global"
	"blog_server/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type OperateLog struct {
	c     *gin.Context
	title string
	level enum.LogLevel
}

func (o *OperateLog) SetTitle(title string) {
	o.title = title
}

func (o *OperateLog) SetLevel(level enum.LogLevel) {
	o.level = level
}

func NewOperateLog(c *gin.Context) *OperateLog {
	return &OperateLog{
		c: c,
	}
}

func (o *OperateLog) Save() {
	ip := o.c.ClientIP()
	addr := core.GetIpAddr(ip)
	userId := uint(1)

	error := global.DB.Create(&models.LogModel{
		LogType: enum.OperateLogType,
		Title:   o.title,
		Content: "",
		Level:   o.level,
		UserID:  userId,
		IP:      ip,
		Addr:    addr,
	}).Error

	if error != nil {
		logrus.Warnf("保存操作日志失败: %v", error)
		return
	}
}
