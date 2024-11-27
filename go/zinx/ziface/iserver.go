package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	//给当前服务注册一个路由业务方法，共客户端连接处理使用
	AddRouter(msgId uint32, Router IRouter)
	GetConnMgr() IConnManager
	//设置该server在连接创建时的Hook函数
	SetOnConnStart(func (IConnection))
	//设置该server在连接断开时的Hook函数
	SetOnConnStop(func (IConnection))
	//调用OnConnStart Hook函数
	CallOnConnStart(conn IConnection)
	//调用OnConnStop Hook函数
	CallOnConnStop(conn IConnection)
}