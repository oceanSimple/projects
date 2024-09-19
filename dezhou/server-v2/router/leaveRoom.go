package router

import (
	"github.com/gin-gonic/gin"
	"server-v2/global"
)

type leaveRoomRequestParams struct {
	RoomNumber int    `json:"roomNumber"`
	Account    string `json:"account"`
}

func leaveRoom(e *gin.Engine) {
	e.POST("/leaveRoom", func(context *gin.Context) {
		// 获取参数
		var params leaveRoomRequestParams
		if err := context.BindJSON(&params); err != nil {
			context.JSON(400, NewResponse(500, "参数错误", nil))
			return
		}
		// 检查房间号是否合法
		if global.Rooms.RoomNumberIsValid(params.RoomNumber) == false {
			context.JSON(400, NewResponse(501, "房间号不存在", nil))
			return
		}

		var room = global.Rooms.GetRoomFromRoomList(params.RoomNumber)
		// 检查用户是否在房间中
		player, exist := room.Players[params.Account]
		if exist == false {
			context.JSON(400, NewResponse(502, "用户不在房间中", nil))
			return
		}
		// 删除用户
		delete(room.Players, params.Account)
		room.PlayerCount--
		// 如果用户是房主, 则随机分配一个房主
		if player.IsHost == 1 {
			for _, p := range room.Players {
				p.IsHost = 1
				break
			}
		}
		context.JSON(200, NewResponse(200, "离开房间成功", nil))
	})
}
