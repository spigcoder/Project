package logService

import (
	"blog_server/core"
	"blog_server/enum"
	"blog_server/global"
	"blog_server/models"

	"github.com/gin-gonic/gin"
)

func SuccessLoginLog(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	//TODO: 根据jwt获得userid
	userID := uint(1)
	//TODO: 根据userid获得username
	userName := ""

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     "",
		UserID:      userID,
		IP:          ip,
		Addr:        addr,
		LoginStatus: true,
		UserName:    userName,
		Pwd:         "",
		LoginType:   loginType,
	})
}

func NewLoginFail(c *gin.Context, loginType enum.LoginType, msg string, username string, pwd string) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录失败",
		Content:     msg,
		IP:          ip,
		Addr:        addr,
		LoginStatus: false,
		UserName:    username,
		Pwd:         pwd,
		LoginType:   loginType,
	})
}
