package handler

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"server-v2/message"
)

type handler func(data gjson.Result) []byte

var handlers = make(map[string]handler)

func init() {
	handlers["/prepare"] = prepare
}

func HandleMessage(str []byte) []byte {
	// 将接收到的消息转换成Message结构体
	var msg message.Message
	err := json.Unmarshal(str, &msg)
	if err != nil {
		return message.GetErrorMessageJson("Invalid message")
	}

	handle, exist := handlers[msg.Path]
	if !exist {
		return message.GetErrorMessageJson("Path not found")
	}

	return handle(gjson.Get(string(str), "data"))
}
