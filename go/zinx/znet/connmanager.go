package znet

import(
	"fmt"
	"sync"
	"errors"
	"zinx/ziface"
)

type ConnManager struct {
	connections	map[uint32] ziface.IConnection
	//全局变量要使用互斥锁来保证数据的正确性
	connLock 	sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32] ziface.IConnection),
	}
}

func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	//加锁来保护共享资源
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	connMgr.connections[conn.GetConnID()] = conn
	fmt.Println("connection add to ConnManager successfully: conn num = ", connMgr.Len())
}

func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	
	delete(connMgr.connections, conn.GetConnID())
	fmt.Println("connection remove ConnID", conn.GetConnID(), " successfully: conn num = ", connMgr.Len())
}

func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}


func (connMgr *ConnManager) ClearConn(){
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	for connID, conn := range connMgr.connections {
		conn.Stop()
		delete(connMgr.connections, connID)
	}
	fmt.Println("Clear All Connection sucessfully: conn num = ", connMgr.Len())
}

func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}