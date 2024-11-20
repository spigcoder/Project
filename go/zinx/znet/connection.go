package znet

import(
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	//当前连接的TCP socket套接字
	Conn *net.TCPConn
	//当前连接的ID，也成为SessionID，这个ID全局唯一
	ConnID uint32
	isClosed bool

	//该连接的处理方法的API
	handleAPI ziface.HandFunc

	// 告知该连接已经退出了
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callBack ziface.HandFunc) *Connection {
	return &Connection{
		Conn: conn,
		ConnID: connID,
		handleAPI: callBack,
		isClosed: false,
		ExitBuffChan: make(chan bool, 1),
	}
}

func (c *Connection) StartRead() {
	fmt.Println("Reader is Running")
	defer fmt.Println(c.RemoteAddr().String(), "conn read exit")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			c.ExitBuffChan <- true
			continue
		}
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID", c.ConnID, "handle is error")
			c.ExitBuffChan <- true
			return
		}
	}
}

func (c *Connection) Start() {
	go c.StartRead()

	for {
		select {
		case <- c.ExitBuffChan:
			return 
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.Conn.Close()
	c.ExitBuffChan <- true
	close(c.ExitBuffChan)
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