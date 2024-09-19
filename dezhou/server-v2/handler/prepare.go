package handler

import (
	"github.com/tidwall/gjson"
	"server-v2/global"
	"server-v2/message"
)

type prepareParams struct {
	Account    string `json:"account"`
	RoomNumber int    `json:"roomNumber"`
	Flag       int    `json:"flag"`
}

var prepare handler = func(data gjson.Result) []byte {
	// 获取参数
	var params = prepareParams{
		Account:    data.Get("account").String(),
		RoomNumber: int(data.Get("roomNumber").Int()),
		Flag:       int(data.Get("flag").Int()),
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
	// 更改玩家准备状态
	player.IsPrepared = params.Flag

	// TODO prepare 通知其他玩家

	return message.GetSuccessMessageJson("success to update player's prepare status")
}
