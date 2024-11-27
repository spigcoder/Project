package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	//当前Connection属于哪个server
	TcpServer	ziface.IServer
	//当前连接的TCP socket套接字
	Conn *net.TCPConn
	//当前连接的ID，也成为SessionID，这个ID全局唯一
	ConnID   uint32
	isClosed bool
	//该连接的处理方法router
	MsgHandle ziface.IMsgHandle
	// 告知该连接已经退出了
	ExitBuffChan chan bool
	//无缓冲通道，用于两个Goroutine之间进行消息通信
	msgChan	  chan []byte
	//有缓冲通道，用于两个Goroutine之间进行消息通信
	msgBuffChan	chan []byte
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandle ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		ExitBuffChan: make(chan bool, 1),
		MsgHandle:    msgHandle,
		msgChan: make(chan []byte),
		msgBuffChan: make(chan []byte, utils.GlobalObject.MaxMsgChanLen),
		TcpServer: server,
	}
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

func (c *Connection) SatrtWriter() {
	fmt.Println("[Writer Goroutine is Running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
			case data := <-c.msgChan:
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Data error: ", err, " conn Writer exit!")
					return
				}
			case data, ok := <-c.msgBuffChan:
				if ok {
					if _, err := c.Conn.Write(data); err != nil {
						fmt.Println("Send Buff Data error: ", err)
						return
					} else {
						break
					}
				}
			case <- c.ExitBuffChan:
				return
		}
	}
}

func (c *Connection) StartRead() {
	fmt.Println("[Reader is Running]")
	defer fmt.Println(c.RemoteAddr().String(), "conn read exit")
	defer c.Stop()

	for {
		dp := NewDataPack()
		//进行数据的读取
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("read msg head err: ", err)
			c.ExitBuffChan <- true
			continue
		}
		//进行拆包操作
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error: ", err)
			c.ExitBuffChan <- true
			continue
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data err: ", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)
		req := Request{
			conn: c,
			msg:  msg,
		}
		if utils.GlobalObject.MaxWorkerPoolSize > 0 {
			c.MsgHandle.SendMsgToTaskQueue(&req)
		} else {
			//从路由器Routers中找到注册绑定conn的对应的handle
			//使用Router提供的方法对于Request的数据进行处理操作
			go c.MsgHandle.DoMsgHandler(&req)
		}
	}
}

func (c *Connection) Start() {
	go c.StartRead()
	go c.SatrtWriter()

	c.TcpServer.CallOnConnStart(c)

	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.TcpServer.CallOnConnStop(c)

	c.Conn.Close()
	c.ExitBuffChan <- true
	c.TcpServer.GetConnMgr().Remove(c)
	close(c.ExitBuffChan)
	close(c.msgBuffChan)
	close(c.msgChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg是一个进行数据发送的过程，这个方法可以让我们发送数据的过程透明化，这样可读性会更好
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.msgChan <- msg

	return nil
}

func (c *Connection) SendBuffMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.msgBuffChan <- msg

	return nil
}
