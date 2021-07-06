package router

// todo Admin路由
//func LoadAdmin(e *gin.Engine) {
//	//需要Token
//	apiNeedToken := e.Group("/v1/api")
//	apiNeedToken.Use(middleware.TokenAuthentication())
//	{
//		user := apiNeedToken.Group("/user")
//		{
//			user.GET("/UserGet", controllers.UserGet) //通过ID查询用户信息
//		}
//	}
//
//	//无需Token
//	apiNOToken := e.Group("/v1/api")
//	{
//		user := apiNOToken.Group("/user")
//		{
//			user.GET("/Login", controllers.Login)        //登录
//			user.POST("/Register", controllers.Register) //注册
//			user.GET("/Test", controllers.Test)          //测试
//		}
//	}
//}
