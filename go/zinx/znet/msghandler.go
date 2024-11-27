package znet

import (
	"fmt"
	"strconv"
	"zinx/utils"
	"zinx/ziface"
)

type MsgHandle struct {
	//MsgHandler中的Apis绑定的就是不同的MsgId对应的路由处理方法
	Apis map[uint32] ziface.IRouter
	WorkerPoolSize	 uint32
	TaskQueue		 []chan ziface.IRequest
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.MaxWorkerPoolSize,
		TaskQueue: make([]chan ziface.IRequest, utils.GlobalObject.MaxWorkerPoolSize),
	}
}

//就是开启一个worker进行工作
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started ")
	for {
		select {
			case request := <- taskQueue:
				mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//每个worker的request缓冲区的大小位WorkerTaskLen
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//启动当前worker，阻塞地等待对应的任务队列是否有消息传来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//根据request的conn来确定由哪个worker进行负责
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(), " request msgID = ", request.GetMsgId(), "to workerID = ", workerID)
	mh.TaskQueue[workerID] <- request
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