package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-v2/model"
)

type loginRequestParams struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func login(e *gin.Engine) {
	e.POST("/login", func(c *gin.Context) {
		var params loginRequestParams
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusOK, NewResponse(400, "参数错误", nil))
			return
		}
		var checked, user = checkAccount(params.Account, params.Password)
		if checked {
			c.JSON(http.StatusOK, NewResponse(200, "登录成功", user))
		} else {
			c.JSON(http.StatusOK, NewResponse(400, "登录失败", nil))
		}
	})
}

func checkAccount(account, password string) (bool, model.User) {
	return true, model.User{}
}
