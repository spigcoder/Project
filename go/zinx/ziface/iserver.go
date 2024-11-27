package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	//给当前服务注册一个路由业务方法，共客户端连接处理使用
	AddRouter(msgId uint32, Router IRouter)
}