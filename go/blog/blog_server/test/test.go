package main

import (
	"blog_server/core"
	"blog_server/flags"
	"blog_server/global"
	"blog_server/utils/jwt"
)

// 测试jwt
// func main() {
// 	flags.Parse()
// 	global.Config = core.ReadConf()
// 	jj, _:= jwt.GetToken(jwt.Claims{
// 		UserID:   123,
// 		UserName: "admin",
// 		Role:     2,
// 	})
// 	x, _ := jwt.ParseToken(jj)
// 	fmt.Println(x)
// }

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()
	token, _ := jwt.GetToken(jwt.Claims{
		UserID:   123,
		UserName: "admin",
		Role:     2,
	})
	jwt.LoseEfficacy(token, jwt.AdminLoseEfficacy)
	res, ok := jwt.IsBlackList(token)
	if ok {
		println(res)
	}
}
