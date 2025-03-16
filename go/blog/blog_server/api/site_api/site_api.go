package siteApi

import (
	res "blog_server/common/res"
	"blog_server/enum"
	logService "blog_server/service/log_service"
	"time"

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
	IP   string `json:"ip"`
}

func (SiteApi) SiteOperationView(c *gin.Context) {
	operLog := logService.GetLogByGin(c)
	// 首先由中间件来进行数据的获取，这里要将请求体中的数据拿出
	operLog.ShowRequest()
	operLog.ShowResponse()
	operLog.SetTitle("更新")
	operLog.SetInfoItem("请求时间", time.Now())

	var cr SiteOperator
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	operLog.SetInfoItem("struct", cr)
	operLog.SetErrItem("Slice", []int{1, 2, 3})
	operLog.SetWarnItem("Array", [2]int{1, 2})
}
