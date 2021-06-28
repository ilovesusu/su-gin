package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/suerror"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//封装返回json数据
func ReturnJson(c *gin.Context, g suerror.GlobeError, data interface{}) {
	var status int
	switch g.Code {
	case 40001:
		status = http.StatusUnprocessableEntity
	case 200:
		status = http.StatusOK
	default:
		status = http.StatusOK
	}
	c.JSON(status, &Response{
		Code: g.Code,
		Msg:  g.Msg,
		Data: data,
	})
}
