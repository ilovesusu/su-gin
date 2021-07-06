package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/controllers"
)

func LoadHioshop(e *gin.Engine) {
	//		/admin
	admin := e.Group("/admin")
	{	//		/admin/index
		admin.GET("/ad", controllers.GetAd)
		admin.GET("/admin", controllers.GetAdmin)
		admin.GET("/category", controllers.GetCategory)
		admin.GET("/goods", controllers.GetGoods)
		admin.GET("/index", controllers.GetIndex)
		admin.GET("/notice", controllers.GetNotice)
		admin.GET("/order", controllers.GetOrder)
		admin.GET("/user", controllers.GetUser)
		admin.GET("/specification", controllers.GetSpecification)
		admin.GET("/shipper", controllers.GetShipper)
		admin.GET("/shopcart",controllers.GetShopcart)
		a2 := admin.Group("/admin")//这里的 a2 是/admin/admin 并不是写重复了
		{
			a2.GET("/showset", controllers.GetShowSet)
		}
		admin.Group("/auth")
		{
			admin.GET("/login",controllers.GetLogin)
		}
		category := admin.Group("/category")
		{
			category.GET("/topCategory",controllers.GetTopCategory)
		}
		goods := admin.Group("/goods")
		{
			goods.Any("/onsale", controllers.GetOnsale)
			goods.GET("/out", controllers.GetOut)
			goods.GET("/drop", controllers.GetDrop)
			goods.GET("/sort",controllers.GetSort)
			//goods.GET("/getAllCategory",controllers.GetGetAllCategory)
			//goods.GET("/getExpressData",controllers.GetGetExpressData)
			//goods.GET("/getAllSpecification",controllers.GetGetAllSpecification)
		}
		index := admin.Group("/index")
		{
			index.GET("/main", controllers.GetMain)
			index.GET("/getQiniuToken",controllers.GetGetQiniuToken)
		}
		shipper := admin.Group("/shipper")
		{
			shipper.GET("/freight", controllers.GetFreight)
			shipper.GET("/list",controllers.GetList)
			shipper.GET("/getareadata",controllers.GetGetareadata)
			shipper.GET("/info",controllers.GetInfo)
		}

		order := admin.Group("/order")
		{
			order.GET("/getAllRegion", controllers.GetGetAllRegion)
			order.GET("/getAutoStatus",controllers.GetGetAutoStatus)

		}
	}
	//需要Token 注册局部中间件
	//apiNeedToken := e.Group("/v1/api")
	//apiNeedToken.Use(middleware.TokenAuthentication())
	//{
	//	user := apiNeedToken.Group("/user")
	//	{
	//		user.GET("/UserGet", controllers.UserGet) //通过ID查询用户信息
	//	}
	//}
	//
	////无需Token
	//apiNOToken := e.Group("/v1/api")
	//{
	//	user := apiNOToken.Group("/user")
	//	{
	//		user.GET("/Login", controllers.Login)        //登录
	//		user.POST("/Register", controllers.Register) //注册
	//		user.GET("/Test", controllers.Test)          //测试
	//	}
	//}
}
