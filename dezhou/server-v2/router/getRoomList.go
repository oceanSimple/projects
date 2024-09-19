package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-v2/global"
	"strconv"
)

func getRoomList(e *gin.Engine) {
	e.GET("/getRoomList", func(c *gin.Context) {
		const pageCount = 10 // 每页显示的房间数

		// 获取page参数
		var pageParam = c.Query("page")
		var page, pageTransferErr = strconv.Atoi(pageParam)
		var msg string
		var msgCode = 200 // 默认为200, 表示成功
		// 如果page参数不为int类型, 或者page小于1, 或者page大于总页数, 均默认使用1
		if pageTransferErr != nil || page < 1 || page > (global.Rooms.MaxRoomNumber/pageCount) {
			page = 1
			msg = "page参数不合法"
			msgCode = 201
		}
		c.JSON(http.StatusOK, NewResponse(msgCode, msg,
			global.Rooms.RoomList[(page-1)*pageCount:page*pageCount]))
	})
}
