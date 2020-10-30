package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//封装返回json数据
func ReturnJson(c *gin.Context, msg string, code int, data []byte) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  msg,
			"code": code,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"code": code,
		"data": string(data),
	})
}

func SuperPrint(super interface{}) {
	fmt.Println("**************************")
	fmt.Println("*\t", super, "\t*")
	fmt.Println("**************************")
}
