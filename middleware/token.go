package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/config"
	"github.com/ilovesusu/su-gin/utils"
)

func TokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Token")
		if tokenStr == "" {
			utils.ReturnJson(c, "Token验证失败", 40002, nil)
			c.Abort()
			return
		} else {
			//下方编辑验证操作
			newjwt:= utils.JWT{}
			sujwt := newjwt.NewJWT()
			claims, err := sujwt.ParseToken(tokenStr)
			fmt.Println(claims)
			if err != nil {
				utils.ReturnJson(c, "Token验证失败"+err.Error(), 40002, nil)
				c.Abort()
				return
			}  else if config.SuApp.Run_Mode == "debug" {

			}
		}
	}
}
