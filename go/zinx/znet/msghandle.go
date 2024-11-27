package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
)

type MsgHandle struct {
	//MsgHandler中的Apis绑定的就是不同的MsgId对应的路由处理方法
	Apis map[uint32] ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgId(), "is not FOUND")
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api, msg = " + strconv.Itoa(int(msgId)))
	}
	mh.Apis[msgId] = router
}