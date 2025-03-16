package router

import (
	"blog_server/global"

	"github.com/gin-gonic/gin"
)

func Run() {
	addr := global.Config.System.Addr()
	gin.SetMode(global.Config.System.GinMode)
	r := gin.Default()
	r.Static("uploads", "uploads")

	api := r.Group("/api")

	SiteRouter(api)
	LogRouter(api)
	r.Run(addr)
}
