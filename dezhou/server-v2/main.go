package main

import (
	"github.com/gin-gonic/gin"
	"server-v2/router"
	"server-v2/ws"
	"strconv"
)

const (
	port = 9000
)

func main() {
	r := gin.Default()
	r.GET("/ws", ws.BuildWsConn)
	router.InitRouter(r)
	r.Run(":" + strconv.Itoa(port))
}
