package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
}

//所有conn在处理业务的函数接口，第一个参数是Socket原生连接，第二个是请求的数据，第三个是数据的长度
type HandFunc func(*net.TCPConn, []byte, int) error