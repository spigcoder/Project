package router

import (
	"blog_server/global"

	"github.com/gin-gonic/gin"
)

func Run() {
	addr := global.Config.System.Addr()
	r := gin.Default()

	api := r.Group("/api")

	SiteRouter(api)
	r.Run(addr)
}
