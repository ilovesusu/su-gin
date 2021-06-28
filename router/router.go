package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/configs"
	"github.com/ilovesusu/su-gin/middleware"
	"io"
	"net/http"
	"os"
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
	// 发行模式
	gin.SetMode(configs.SuApp.Run_Mode)
	// 如果是api则记录到文件,否则直接打印到控制台
	if configs.SuApp.SaveLog == true {
		f, _ := os.OpenFile(configs.SuApp.Log+"/gin_access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		f1, _ := os.OpenFile(configs.SuApp.Log+"/gin_error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		//为了方便查阅,控制台与文件同时输出(io reader writer 会增加负载)
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		gin.DefaultErrorWriter = io.MultiWriter(f1, os.Stdout)
	}
	r := gin.Default()

	// 注册全局中间件
	r.Use(middleware.PageNotFound())
	r.Use(middleware.Cors())

	r.StaticFS("/assets", http.Dir(configs.SuApp.Assets)) //静态资源

	LoadUser(r)
	LoadAdmin(r)

	return r
}
