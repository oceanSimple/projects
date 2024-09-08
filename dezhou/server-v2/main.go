package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"server-v2/model"
	"server-v2/output"
	"strconv"
)

const (
	port = 9000
)

var (
	playerList = make(map[string]*model.Player) // key is ip
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	r.GET("/ws", buildWsConn)
	r.Run(":" + strconv.Itoa(port))
}

func buildWsConn(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(output.Error() + err.Error())
		return
	}
	log.Println(output.Default()+"client connect:", conn.RemoteAddr())
	defer conn.Close()

	// Create a player
	player := &model.Player{
		Ip:   conn.RemoteAddr().String(),
		Conn: conn,
	}
	playerList[player.Ip] = player

	// Handle WebSocket connection
	for {
		// Read message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(output.Default() + err.Error())
			return
		}
		log.Println("Received message:", string(p))

		// Send message
		err = conn.WriteMessage(messageType, []byte("Hello, world!"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
