package model

import (
	"server-v2/card"
)

type Room struct {
	Number  int                // Room number
	Status  int                // 0: no gaming, 1: gaming
	Players map[string]*Player // player map, key is ip

	CardBox *card.CardBox // card box
}

func NewRoom(number int) *Room {
	return &Room{
		Number:  number,
		Status:  0,
		Players: make(map[string]*Player),

		CardBox: card.NewCardBox(),
	}
}
