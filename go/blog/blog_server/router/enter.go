package router

import (
	"blog_server/global"
	"blog_server/middle"

	"github.com/gin-gonic/gin"
)

func Run() {
	addr := global.Config.System.Addr()
	r := gin.Default()

	api := r.Group("/api")
	api.Use(middle.OperatorLogMiddle)

	SiteRouter(api)
	r.Run(addr)
}
