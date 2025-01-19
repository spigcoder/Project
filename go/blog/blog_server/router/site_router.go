package router

import (
	"blog_server/api"

	"github.com/gin-gonic/gin"
)

func SiteRouter(g *gin.RouterGroup) {
	app := api.App.SiteApi

	g.GET("/site", app.SiteInfoView)
}
