package router

import (
	"github.com/gin-gonic/gin"
	"server-v2/global"
	"server-v2/model"
)

type enterRoomRequestParams struct {
	RoomNumber int    `json:"roomNumber"`
	Account    string `json:"account"`
}

func enterRoom(e *gin.Engine) {
	e.POST("/enterRoom", func(c *gin.Context) {
		// 读取参数
		var params enterRoomRequestParams
		if err := c.BindJSON(&params); err != nil {
			c.JSON(200, NewResponse(201, err.Error(), nil))
			return
		}
		// 判断房间号是否合法
		var roomNumberIsValid = global.Rooms.RoomNumberIsValid(params.RoomNumber)
		if !roomNumberIsValid {
			c.JSON(200, NewResponse(202, "房间号错误！", nil))
			return
		}

		// 判断房间是否已满
		var room = global.Rooms.GetRoomFromRoomList(params.RoomNumber)
		var roomIsFull = room.IsFull()
		if roomIsFull {
			c.JSON(200, NewResponse(203, "房间已满！", nil))
			return
		}

		// 加入房间
		// 判断玩家是否是房间内唯一人, 授予房主权限
		var isHost = 0
		var isPrepared = 0
		if room.PlayerCount == 0 {
			isHost = 1
			isPrepared = 1
		}
		// 更改房间信息
		room.PlayerCount++
		room.Players[params.Account] = &model.RoomPlayer{
			Account:    params.Account,
			Chip:       0,
			IsPrepared: isPrepared,
			IsHost:     isHost,
		}

		// 更改玩家信息
		//var player = global.Players.GetPlayerFromPlayerList(params.Account)
		//player.RoomNumber = params.RoomNumber

		// 获取房间内玩家的信息
		var roomPlayers = global.Rooms.GetRoomFromRoomList(params.RoomNumber).Players
		c.JSON(200, NewResponse(200, "加入房间成功！", roomPlayers))
	})
}
