package main

import (
	router "blog_server/Router"
	"blog_server/core"
	"blog_server/flags"
	"blog_server/global"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	core.InitIP()
	flags.Run()

	router.Run()
}
