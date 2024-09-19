package router

import (
	"github.com/gin-gonic/gin"
	"server-v2/global"
)

// 转让房主权限

type transferHostRequestParams struct {
	RoomNumber int    `json:"roomNumber"` // 房间号
	OriginHost string `json:"originHost"` // 原房主, 也就是发起请求的用户
	TargetHost string `json:"targetHost"` // 目标房主, 也就是要转让给的用户
}

func transferHost(e *gin.Engine) {
	e.POST("/transferHost", func(c *gin.Context) {
		// 读取参数
		var params transferHostRequestParams
		if err := c.BindJSON(&params); err != nil {
			c.JSON(200, NewResponse(400, err.Error(), nil))
			return
		}

		// 判断房间号是否合法
		if global.Rooms.RoomNumberIsValid(params.RoomNumber) == false {
			c.JSON(200, NewResponse(400, "房间号错误！", nil))
			return
		}

		var room = global.Rooms.GetRoomFromRoomList(params.RoomNumber)
		// 判断originHost是否是房主
		originHostPlayer := room.Players[params.OriginHost]
		if originHostPlayer == nil || originHostPlayer.IsHost == 0 {
			c.JSON(200, NewResponse(400, "您不是房主！", nil))
			return
		}
		// 判断targetHost是否在房间内
		targetHostPlayer := room.Players[params.TargetHost]
		if targetHostPlayer == nil {
			c.JSON(200, NewResponse(400, "目标玩家不在房间内！", nil))
			return
		}
		// 转让房主权限
		originHostPlayer.IsHost = 0
		targetHostPlayer.IsHost = 1

		c.JSON(200, NewResponse(200, "房主转让成功！", nil))
	})
}
