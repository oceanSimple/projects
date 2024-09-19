package handler

import (
	"github.com/tidwall/gjson"
	"server-v2/global"
	"server-v2/message"
)

type startGameParams struct {
	Account    string `json:"account"`
	RoomNumber int    `json:"roomNumber"`
}

var startGame handler = func(data gjson.Result) []byte {
	// 获取参数
	var params = startGameParams{
		Account:    data.Get("account").String(),
		RoomNumber: int(data.Get("roomNumber").Int()),
	}
	// 获取房间对象
	room := global.Rooms.GetRoomFromRoomList(params.RoomNumber)
	if room == nil {
		return message.GetErrorMessageJson("Room not found")
	}
	// 获取玩家对象
	player := room.Players[params.Account]
	if player == nil {
		return message.GetErrorMessageJson("Player not found")
	}

	// 查看玩家是否是房主
	if player.IsHost != 1 {
		return message.GetErrorMessageJson("You are not the host")
	}
	// 查看是否所有玩家都准备好了
	for _, player := range room.Players {
		if player.IsPrepared == 0 {
			return message.GetErrorMessageJson("Not all players are prepared")
		}
	}
	// TODO startGame 通知其他玩家, 开始游戏

	return message.GetSuccessMessageJson("success to start game")
}
