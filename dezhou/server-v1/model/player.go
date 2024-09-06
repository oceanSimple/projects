package model

import "github.com/gorilla/websocket"

type Player struct {
	Conn       *websocket.Conn
	Ip         string
	Nickname   string
	RoomNumber int // 0 stands for no room, 1~n stands for room number
}
