package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/configs"
	"github.com/ilovesusu/su-gin/suerror"
	"github.com/ilovesusu/su-gin/utils"
)

func TokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Token")
		if tokenStr == "" {
			utils.ReturnJson(c,suerror.SIGNERROR, nil)
			c.Abort()
			return
		} else {
			//下方编辑验证操作
			//newjwt:= utils.JWT{}
			//sujwt := newjwt.NewJWT()
			claims, err := utils.ParseToken(tokenStr)
			fmt.Println(claims)
			if err != nil {
				utils.ReturnJson(c, suerror.SIGNERROR, nil)
				c.Abort()
				return
			}  else if configs.SuApp.Run_Mode == "debug" {

			}
		}
	}
}
