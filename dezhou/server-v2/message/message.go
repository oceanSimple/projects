package message

import "encoding/json"

type Message struct {
	// 请求路径
	Path string `json:"path"`
	// 消息数据
	Data any `json:"data"`
}

func GetMessageJson(message Message) []byte {
	marshal, err := json.Marshal(message)
	if err != nil {
		return nil
	} else {
		return marshal
	}
}

func GetErrorMessageJson(errMsg string) []byte {
	msg := Message{
		Path: "/error",
		Data: errMsg,
	}
	return GetMessageJson(msg)
}

func GetSuccessMessageJson(str string) []byte {
	msg := Message{
		Path: "/success",
		Data: str,
	}
	return GetMessageJson(msg)
}
