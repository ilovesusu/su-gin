package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/config"
	"github.com/ilovesusu/su-gin/controllers"
	"github.com/ilovesusu/su-gin/middleware"
	"net/http"
)

// post delete put get 对应 增 删 改 查

/*
使用接收单个参数各种方法：
c.Query            示例: welcome?firstname=Jane&lastname=Doe  可以匹配 firstname
c.DefaultQuery     示例: 同上,可以设置默认值
c.Param            示例: "/user/:name"  可以匹配:name c.Query
c.PostForm         示例: 无论是 multipart/form-data格式还是application/x-www-form-urlencoded格式,都可以获取到参数。
c.DefaultPostForm  示例: 同上,可以设置默认值
c.QueryMap         示例: 从url中获取map   传参xxx?ids[a]=1234&ids[b]=hello
c.PostFormMap      示例: 从Body中获取map 传参 ids[0]:234561
c.FormFile         示例: 接受上传文件
c.MultipartForm    示例: 处理多个文件的上传
*/

func InitRouter() *gin.Engine {
	//发行模式
	gin.SetMode(config.SuApp.Run_Mode)
	r := gin.Default()

	r.StaticFS("/assets", http.Dir(config.SuApp.Assets)) //静态资源

	//需要Token
	apiNeedToken := r.Group("/v1/api")
	apiNeedToken.Use(middleware.TokenAuthentication())
	{
		user := apiNeedToken.Group("/user")
		{
			user.GET("/UserGet", controllers.UserGet) //通过ID查询用户信息
		}
	}

	//无需Token
	apiNOToken := r.Group("/v1/api")
	{
		user := apiNOToken.Group("/user")
		{
			user.GET("/Login", controllers.Login)        //登录
			user.POST("/Register", controllers.Register) //注册
			user.GET("/Test", controllers.Test)          //测试
		}
	}
	return r
}
