package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
	"zinx/utils"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	msgHandler ziface.IMsgHandle
	ConnMgr	  ziface.IConnManager
	OnConnStart func(conn ziface.IConnection)
	OnConnStop func(conn ziface.IConnection)
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

func (s *Server) SetOnConnStart(hook func (ziface.IConnection)) {
	s.OnConnStart = hook
}

func (s *Server) SetOnConnStop(hook func (ziface.IConnection)) {
	s.OnConnStop = hook
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart...")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStart...")
		s.OnConnStop(conn)
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listerner at IP: %s, Port: %d, is starting\n", s.IP, s.Port)
	// 开启一个goroutine去做服务端listen业务
	go func() {
		s.msgHandler.StartWorkerPool()
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
			//如果超过最大连接数则关闭此连接
			if s.ConnMgr.Len() > utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}
			dealConn := NewConnection(s, conn, cid, s.msgHandler)
			cid++
			//启动当前连接处理业务
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name: ", s.Name)
	s.ConnMgr.ClearConn()
}

func (s *Server) Server() {
	s.Start()

	// 阻塞，否则主Go退出，listener的go将会退出
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router secc! ")
}

func (s *Server)GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

// 返回的是一个接口，server实现了IServer中的所有的方法，就实现了这个接口
func NewServer() ziface.IServer {
	//使用全局配置文件进行server的配置
	utils.GlobalObject.Reload()
	return &Server{
		Name: utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalObject.Host,
		Port: utils.GlobalObject.TcpPort,
		msgHandler: NewMsgHandle(),
		ConnMgr: NewConnManager(),
	}
}


