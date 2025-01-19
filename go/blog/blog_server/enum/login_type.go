package enum

type LoginType int8

const (
	UserPwdLoginType LoginType = iota + 1 // 用户名密码登录
	QQLoginType                           // QQ号登录
	EmailLoginType                        // 邮箱登录
)
