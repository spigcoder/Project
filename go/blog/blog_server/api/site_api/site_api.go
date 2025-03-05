package siteApi

import (
	"blog_server/enum"
	logService "blog_server/service/log_service"

	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	logService.FailLogin(c, enum.UserPwdLoginType, "密码错误", "spigcoder", "123456")
	logService.SuccessLogin(c, enum.UserPwdLoginType)
}

type SiteOperator struct {
	Name string `json:"name"`
}

func (SiteApi) SiteOperationView(c *gin.Context) {
	operLog := logService.GetLogByGin(c)
	// 首先由中间件来进行数据的获取，这里要将请求体中的数据拿出
	operLog.ShowRequest()
	operLog.ShowResponse()

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
