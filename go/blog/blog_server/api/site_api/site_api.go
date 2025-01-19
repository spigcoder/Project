package siteApi

import (
	"blog_server/enum"
	logService "blog_server/service/log_service"

	"github.com/gin-gonic/gin"
)


type SiteApi struct {
    
}

func (SiteApi)SiteInfoView(c *gin.Context) {
	logService.FailLogin(c, enum.UserPwdLoginType, "密码错误", "spigcoder", "123456")
	logService.SuccessLogin(c, enum.UserPwdLoginType) 
}
