package znet

import "zinx/ziface"

// request没有长度也没有定义消息类型，这样对于数据的管理是不充分的
type Request struct {
	conn ziface.IConnection
	msg ziface.IMessage
}

func(r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func(r *Request) GetData() []byte {
	return r.msg.GetData()
}

func(r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}