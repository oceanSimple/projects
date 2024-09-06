# 顺子算法
```go
package pattern

import (
	"server-v1/card"
)

func HasStraight(cards []*card.Card) (bool, [][]*card.Card) {
	if len(cards) != 7 {
		return false, nil
	}

	var result = make([][]*card.Card, 0)

	var middleArray = make([][][]*card.Card, 0)
	// common logic
	for i := 0; i < 3; i++ {
		// a temp slice to store the straight cards
		var temp = make([][]*card.Card, 0)

		// add the first card
		// [{n_0}]
		temp = append(temp, []*card.Card{cards[i]})

		// the next card's weight we need to find
		var standard = cards[i].Number.Weight - 1
		var index = i + 1 // the index of the last card in temp

		// find card_1
		// because the length of cards is 7, so card_1 can at most be cards[3]
		index, card_1_Array := getCardNArray(cards, index, 3, standard)
		if len(card_1_Array) == 0 { // did not find card_1
			continue
		} else {
			temp = append(temp, card_1_Array)
			standard--
		}

		// find card_2
		index, card_2_Array := getCardNArray(cards, index, 4, standard)
		if len(card_2_Array) == 0 {
			continue
		} else {
			temp = append(temp, card_2_Array)
			standard--
		}

		// find card_3
		index, card_3_Array := getCardNArray(cards, index, 5, standard)
		if len(card_3_Array) == 0 {
			continue
		} else {
			temp = append(temp, card_3_Array)
			standard--
		}

		// find card_4
		index, card_4_Array := getCardNArray(cards, index, 6, standard)
		if len(card_4_Array) == 0 {
			continue
		} else {
			temp = append(temp, card_4_Array)
			standard--
		}

		// now the length of temp is 5, store it in middleArray
		middleArray = append(middleArray, temp)
	}

	for i := 0; i < len(middleArray); i++ {
		array := queueDfs(middleArray[i])
		result = append(result, array...)
	}

	// TODO: 特殊情况，A2345
	if len(result) > 0 {
		return true, result
	} else {
		return false, nil
	}
}

// return the index of the last index in the array, and the array
func getCardNArray(mainCardArray []*card.Card, start, end int, standard int) (int, []*card.Card) {
	var result = make([]*card.Card, 0)
	var i = start
	for ; i <= end; i++ {
		if mainCardArray[i].Number.Weight == standard {
			result = append(result, mainCardArray[i])
		} else if mainCardArray[i].Number.Weight > standard {
			continue
		} else {
			break
		}
	}
	return i, result
}

func queueDfs(cardArray [][]*card.Card) [][]*card.Card {
	l := make([][]*card.Card, 0, len(cardArray))
	// put the first array into the queue
	l = append(l, cardArray[0])
	for i := 1; i < len(cardArray); i++ {
		var temp = make([][]*card.Card, 0)
		for j := 0; j < len(cardArray[i]); j++ { // suffix
			for n := 0; n < len(l); n++ { // prefix
				temp = append(temp, append(l[n], cardArray[i][j]))
			}
		}
		l = temp
	}
	return l
}

```


# websocket main函数
```go
package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"server-v1/model"
	"server-v1/output"
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
	// Start a http server
	http.HandleFunc("/ws", handleWebSocket)
	log.Println(output.Info() + "Server started on :" + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
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

```
