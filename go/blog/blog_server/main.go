package main

import (
	"blog_server/core"
	"blog_server/flags"
	"blog_server/global"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	core.InitIP()
	fmt.Println(core.GetIpAddr("23.23.4.2"))
}
