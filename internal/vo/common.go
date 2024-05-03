package vo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	ERROR   = 400
	SUCCESS = 0
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func ResultErr(err error, c *gin.Context) {
	var e *Err
	if err != nil {
		if !errors.As(err, &e) {
			c.JSON(http.StatusOK, Response{
				Code: ERROR,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, Response{
			Code: e.code,
			Msg:  e.msg,
		})
		c.Abort()
		return
	}
	c.Abort()
	return
}

func OK(c *gin.Context) {
	Result(SUCCESS, nil, "success", c)
}

func OkData(data, any, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OkMsg(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}
