package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	SendMsg(msgId uint32, data []byte) error
	SendBuffMsg(mdgID uint32, data []byte) error
}

//所有conn在处理业务的函数接口，第一个参数是Socket原生连接，第二个是请求的数据，第三个是数据的长度
//但是这里有一个问题，就是我们呢的业务处理逻辑都是绑死在这个格式当中，那我们可不可以定义一个interface
//然后使用这个interface来让用户填写任意格式的来连接处理业务方法
type HandFunc func(*net.TCPConn, []byte, int) error