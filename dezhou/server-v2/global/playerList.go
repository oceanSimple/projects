package global

import (
	"server-v2/model"
	"sync"
)

type PlayerListClass struct {
	PlayerList map[string]*model.Player
	rwMutex    *sync.RWMutex
}

func newPlayerList() *PlayerListClass {
	return &PlayerListClass{
		PlayerList: make(map[string]*model.Player),
		rwMutex:    &sync.RWMutex{},
	}
}

func (p *PlayerListClass) GetPlayerFromPlayerList(playerID string) *model.Player {
	p.rwMutex.RLock()
	defer p.rwMutex.RUnlock()
	return p.PlayerList[playerID]
}
