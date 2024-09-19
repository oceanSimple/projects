package global

import "server-v2/model"

func initSomeRoom() {
	Rooms.RoomList[18].PlayerCount = 2
	Rooms.RoomList[18].Players["player1"] = &model.RoomPlayer{
		Account:    "play",
		Chip:       0,
		IsPrepared: 0,
		IsHost:     1,
	}
	Rooms.RoomList[18].Players["player2"] = &model.RoomPlayer{
		Account:    "player2",
		Chip:       0,
		IsPrepared: 0,
		IsHost:     0,
	}
}
