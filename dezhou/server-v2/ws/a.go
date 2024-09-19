package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"server-v2/handler"
	"server-v2/output"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func BuildWsConn(c *gin.Context) {
	// 将HTTP连接升级为WebSocket连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(output.Error() + err.Error())
		return
	}
	defer conn.Close()
	defer fmt.Println("连接关闭")
	// 处理消息
	for {
		// 我们默认接收和发送的消息类型都是文本类型
		// 所以messageType返回值我们直接不要了
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(output.Default() + err.Error())
			return
		}

		response := handler.HandleMessage(p)

		// 发送消息
		err = conn.WriteMessage(websocket.TextMessage, response)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
