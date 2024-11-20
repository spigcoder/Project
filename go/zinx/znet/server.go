package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

//这个函数是对已经到来的数据进行处理，cnt是得到的数据的长度，conn是socket连接
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToclient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient err")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Println("[START] Server listerner at IP: %s, Port: %d, is starting\n", s.IP, s.Port)
	// 开启一个goroutine去做服务端listen业务
	go func() {
		//1. 获取一个TCP addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s: %d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		//2. 监听服务器地址
		listen, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen: ", s.IP, "err: ", err)
			return
		}
		//监听成功
		fmt.Println("start Zinx Server: ", s.Name, "succ, now listening... ")
		//这里应该有一个cid的随机生成函数
		var cid uint32
		cid = 0
		//3. 启动server网络连接服务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			//启动当前连接处理业务
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name: ", s.Name)
	//...
}

func (s *Server) Server() {
	s.Start()

	// 阻塞，否则主Go退出，listener的go将会退出
	select {}
}

// 返回的是一个接口，server实现了IServer中的所有的方法，就实现了这个接口
func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
}
