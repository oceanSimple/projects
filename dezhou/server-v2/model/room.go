package model

import (
	"server-v2/card"
)

type Room struct {
	// 房间的号码，唯一标识
	Number int

	// 房间的人数
	PlayerCount int

	// 房间是否已满
	// IsFull int

	// 房间的状态：0-未开始，1-进行中
	Status int

	// 房间的玩家map
	Players map[string]*RoomPlayer `json:"-"`

	// 房间的牌盒
	CardBox *card.CardBox `json:"-"`
}

func NewRoom(number int) *Room {
	return &Room{
		Number:  number,
		Status:  0,
		Players: make(map[string]*RoomPlayer),

		CardBox: card.NewCardBox(),
	}
}

// 判断房间是否已满
func (r *Room) IsFull() bool {
	// (总牌数 - 公共牌数) / 每个玩家受伤的牌数
	return r.PlayerCount >= ((13*4)-5)/2
}

// 判断房间是否空闲
func (r *Room) IsSpare() bool {
	return r.Status == 0
}
