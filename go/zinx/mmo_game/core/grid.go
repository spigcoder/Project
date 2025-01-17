package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID			int				//格子 ID
	MinX		int				//格子左边界坐标
	MaxX		int		
	MinY		int				//格子上边界坐标
	MaxY		int
	playerIDs	map[int]bool	//当前格子内玩家或物体成员ID
	pIDLock		sync.RWMutex
}

func NewGrid(gID, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID: gID,
		MinX: minX,
		MaxX: maxX,
		MinY: minY,
		MaxY: maxY,
		playerIDs: make(map[int]bool),
	}
}

func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.playerIDs[playerID] = true
}

func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerID)
}

func (g *Grid) GetPlayerIDs() (playersIDs []int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	for k, _ := range g.playerIDs {
		playersIDs = append(playersIDs, k)
	}

	return 
}

//打印信息
func (g *Grid) String() string {
	return fmt.Sprintf("Grid id: %d, minX: %d, maxX: %d, minY: %d, maxY: %d, playersIDs: %v", g.GID,
	g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)	
}