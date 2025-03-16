package router

import (
	"blog_server/api"
	"blog_server/middle"

	"github.com/gin-gonic/gin"
)

func SiteRouter(g *gin.RouterGroup) {
	g.Use(middle.OperatorLogMiddle)
	app := api.App.SiteApi

	g.PUT("/oper", app.SiteOperationView)
}
