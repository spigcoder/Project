package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// Server模块的测试函数
type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("recv from client : msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(0, []byte("ping ping ping...\n"))
	if err != nil {
		fmt.Println("call back ping err: ", err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	fmt.Println("recv from client : msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.6\n"))
	if err != nil {
		fmt.Println("call back ping err: ", err)
	}
}

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectionBegin is Called...!")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionEnd(conn ziface.IConnection) {
	fmt.Println("DoConnectionEnd is Called...!")
}

func main() {
	s := znet.NewServer()
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionEnd)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Server()
}

//这里是对datapack拆包和封包的功能进行测试

// func main() {
// 	//创建socket TCP Server
// 	listener, err := net.Listen("tcp", "127.0.0.1:7777")
// 	if err != nil {
// 		fmt.Println("server listen err: ", err)
// 		return
// 	}

// }

//这是对datapack进行测试操作
// func main() {
// 	listener, err := net.Listen("tcp", "127.0.0.1:7777")
// 	if err != nil {
// 		fmt.Println("server listen err: ", err)
// 		return
// 	}
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("server accept err: ", err)
// 		}
// 		go func (conn net.Conn) {
// 			dp := znet.NewDataPack()
// 			for {
// 				headData := make([]byte, dp.GetHeadLen())
// 				//ReadFull只会读取headData的长度到HeadData当中去，如果没有这个长度会阻塞等待
// 				_, err := io.ReadFull(conn, headData)
// 				if err != nil {
// 					fmt.Println("read head error")
// 					break
// 				}
// 				//返回的msgHead是一个接口，如果要使用要进行类型断言
// 				msgHead, err := dp.Unpack(headData)
// 				if err != nil {
// 					fmt.Println("server unpack err:", err)
// 					return
// 				}
// 				if msgHead.GetDataLen() > 0 {
// 					msg := msgHead.(* znet.Message)
// 					msg.Data = make([]byte, msg.GetDataLen())
// 					_, err := io.ReadFull(conn, msg.Data)
// 					if err != nil {
// 						fmt.Println("server unpack data err: ", err)
// 						return
// 					}
// 					fmt.Println("==> rECV Msg: ID=", msg.Id, "Len=", msg.DataLen, "data=", string(msg.Data))
// 				}
// 			}
// 		}(conn)
// 	}
// }
