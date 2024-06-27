package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

//const (  如果一下是公共的放回方法 两种情况可能满足不了  比如一个人有无权限 未授权 等等
//	ERROR   = false
//	SUCCESS = true
//)

type Response struct {
	State int         `json:"state"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		State: code,
		Data:  data,
		Msg:   msg,
	})
}

func OkWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
