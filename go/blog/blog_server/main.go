package main

import (
	"fmt"
	"blog_server/core"
	"blog_server/flags"
	"blog_server/global"
	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	fmt.Println(global.Config)
	core.InitLogrus()
	logrus.Info("config: ", global.Config)
	logrus.Warnln("xxx")
	logrus.Debugln("yyy")
	logrus.Errorln("zzz")
}
