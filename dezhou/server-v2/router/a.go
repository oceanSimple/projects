package router

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func InitRouter(e *gin.Engine) {
	getRoomList(e)
	login(e)
	transferHost(e)
	enterRoom(e)
	leaveRoom(e)
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
