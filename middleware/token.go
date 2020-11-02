package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/config"
	"github.com/ilovesusu/su-gin/db"
	"github.com/ilovesusu/su-gin/utils"
	"strconv"
	"time"
)

func TokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Token")
		if tokenStr == "" {
			utils.ReturnJson(c, "Token验证失败", 40002, nil)
			c.Abort()
			return
		} else {
			newjwt:= utils.JWT{}
			sujwt := newjwt.NewJWT()
			claims, err := sujwt.ParseToken(tokenStr)
			if err != nil {
				utils.ReturnJson(c, "Token验证失败"+err.Error(), 40002, nil)
				c.Abort()
				return
			} else if claims.UID != 0 { //如果是刷手token
				ip := c.ClientIP()
				//如果是黑名单token
				get := db.SuRedis.Get("heitoken:" + tokenStr)
				if get.Val() != "" {
					utils.ReturnJson(c, "Token验证失败", 40002, nil)
					c.Abort()
					return
				}
				//获取userid是否是当前ip登录
				itoa := strconv.Itoa(int(claims.UID))
				cmd := db.SuRedis.Get("uip:" + itoa + ":" + ip)
				//如果不存在该键  错误ip
				if cmd.Val() == "" {
					//拉黑该token  重新登录
					set := db.SuRedis.Set("heitoken:"+tokenStr, 1, 24*time.Hour)
					if set.Err() != nil {
						fmt.Println(set.Err())
					}
					utils.ReturnJson(c, "Token验证失败", 40002, nil)
					c.Abort()
					return
				}
				//成功
				c.Set("sutoken",claims.UID)
			} else if claims.AID != 0  {//总后台
				//成功
				c.Set("sutoken",claims.AID)
			} else if claims.MID != 0  {//商家端
				//成功
				c.Set("sutoken",claims.MID)
			} else if config.SuApp.Run_Mode == "debug" {

			}
		}
	}
}
