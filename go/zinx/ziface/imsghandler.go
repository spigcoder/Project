package ziface

//定义这个接口的目的是为了将不同的msgId与Router进行分离开来
//这样就可以绑定多个业务处理方法

type IMsgHandle interface {
	DoMsgHandler (request IRequest)
	AddRouter (msgId uint32, router IRouter)
	StartWorkerPool()
	SendMsgToTaskQueue(request IRequest)
}