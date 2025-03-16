package router

import (
	"blog_server/api"

	"github.com/gin-gonic/gin"
)

func LogRouter(g *gin.RouterGroup) {
	app := api.App.LogApi

	g.GET("/logs", app.LogListView)
}
