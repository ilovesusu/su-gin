package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/db"
	"github.com/ilovesusu/su-gin/model"
	"github.com/ilovesusu/su-gin/suerror"
	"github.com/ilovesusu/su-gin/utils"
	"time"
)

func GetUser(ctx *gin.Context) {
	type param struct {
		Page int `form:"page"`
	}

	type userDate struct {
		Count       int64         `json:"count"`
		CurrentPage int           `json:"currentPage"`
		PageSize    int           `json:"pageSize"`
		TotalPages  int           `json:"totalPages"`
		Users       *[]model.User `json:"data"`
	}
	type Data struct {
		UserData userDate `json:"userData"`
	}
	var suParam param
	if err := ctx.ShouldBind(&suParam); err != nil {
		suParam.Page = 1
		//utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		//return
	}
	fmt.Println(suParam)
	var user model.User
	//var users []model.User
	userDates := userDate{}
	users, _ := user.GetUserWithPage(suParam.Page)

	userDates.Users = &users
	userDates.Count, _ = user.GetUserCountNum()
	userDates.PageSize = 10
	userDates.TotalPages = (len(users) / 10) + 1
	userDates.CurrentPage = suParam.Page

	var data Data
	data.UserData = userDates
	utils.ReturnJson(ctx, suerror.SUCCESS, data)

}

func GetIndex(ctx *gin.Context) {

	type Data struct {
		GoodsOnsale     int64
		OrderToDelivery int64
		Timestamp       int64
		User            *int64
	}

	//var user model.User
	//var users []model.User
	data := Data{}

	var count int64
	data.User = &count

	data.Timestamp = time.Now().Unix()

	//售卖中
	//var order model.Order
	var goods model.Goods
	//直接操作了Model依赖
	data.GoodsOnsale, _ = goods.GetGoodsOnsaleNumber()

	//db.SuDB.Debug().Model(&model.Goods{}).Count(&OnSale)
	fmt.Println("售卖中：", data.GoodsOnsale)

	db.SuDB.Debug().Model(&model.User{}).Count(&count)

	fmt.Println(count)
	db.SuDB.Debug().Model(&model.Order{}).Where("order_status", 300).Count(&data.OrderToDelivery)
	fmt.Println("未送达", data.OrderToDelivery)
	utils.ReturnJson(ctx, suerror.SUCCESS, data)

}
func GetLogin(ctx *gin.Context) {

}
func GetShowSet(ctx *gin.Context) {
	settings := model.ShowSettings{}
	err := settings.GetShowSettings()
	if err != nil {
		return
	}
	type Data struct {
		Id               int64 `json:"id"`
		Banner           int64 `json:"banner"`
		Channel          int64 `json:"channel"`
		Index_banner_img int64 `json:"index_banner_img"`
		Notice           int64 `json:"notice"`
	}
	data := Data{Id: settings.ID, Banner: settings.Banner, Channel: settings.Channel, Index_banner_img: settings.IndexBannerImg, Notice: settings.Notice}
	fmt.Println("返回前", data)
	//
	utils.ReturnJson(ctx, suerror.SUCCESS, data)

}
func GetAd(ctx *gin.Context) {
	type Data struct {
		Count       int         `json:"count"`
		TotalPages  int         `json:"totalPages"`
		PageSize    int         `json:"pageSize"`
		CurrentPage int         `json:"currentPage"`
		Ads         *[]model.Ad `json:"data"`
	}
	type param struct {
		Page int `form:"page" json:"page" binding:"required"`
	}
	var suParam param
	err := ctx.ShouldBind(&suParam)
	if err != nil {
		return
	}
	var ad model.Ad
	ads, err := ad.GetAdsWithPage(suParam.Page)
	AdsData := Data{}
	AdsData.Ads = &ads
	for i := range ads {
		endTime := ads[i].EndTime
		timeString := time.Unix(endTime, 0).Format("2006-01-02 15:04:05")
		fmt.Println(timeString)
	}
	utils.ReturnJson(ctx, suerror.SUCCESS, AdsData)

}

func GetNotice(ctx *gin.Context) {
	var notice model.Notice

	data, err := notice.GetAllNotices()
	if err != nil {
		return
	}
	utils.ReturnJson(ctx, suerror.SUCCESS, data)

}

func GetShipper(ctx *gin.Context) {
	type Data struct {
		Info []model.Shipper `json:"info"`
		Set  model.Settings  `json:"set"`
	}
	var data Data
	var shipper model.Shipper
	enableShippers, err := shipper.GetAllEnableShippers()
	if err != nil {
		return
	}
	data.Info = enableShippers
	var setting model.Settings
	oneSetting, err := setting.GetOneSetting()
	data.Set = oneSetting
	utils.ReturnJson(ctx, suerror.SUCCESS, data)

}
func GetAdmin(ctx *gin.Context) {
	var admin model.Admin

	a, err := admin.GetAllAdmins()
	if err != nil {
		return
	}
	utils.ReturnJson(ctx, suerror.SUCCESS, a)

}

func GetGoods(ctx *gin.Context) {
	type Param struct {
		Page int    `form:"page"`
		Name string `form:"name"`
	}
	type Date struct {
		Count       int            `json:"count"`
		CurrentPage int            `json:"currentPage"`
		PageSize    int            `json:"pageSize"`
		TotalPages  int            `json:"totalPages"`
		Goods       *[]model.Goods `json:"data"`
	}
	var data Date
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}
	var goods model.Goods
	allGoods, err := goods.GetAllGoods()
	if err != nil {
		return
	}
	data.Count = len(allGoods)
	data.CurrentPage = suParam.Page
	data.PageSize = 10
	data.TotalPages = (len(allGoods) / 10) + 1
	data.Goods = &allGoods

	utils.ReturnJson(ctx, suerror.SUCCESS, data)

}
func GetOrder(ctx *gin.Context) {
	type Param struct {
		Page          int    `form:"page"`
		Logistic_code string `form:"logistic_code"`
		Status        []int  `form:"status"`
	}
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}

}
func GetCategory(ctx *gin.Context) {
	var category model.Category
	allCategory, err := category.GetAllCategory()
	if err != nil {
		return
	}
	utils.ReturnJson(ctx, suerror.SUCCESS, allCategory)
}
func GetTopCategory(ctx *gin.Context) {
	var category model.Category
	allCategory, err := category.GetTopCategory()
	if err != nil {
		return
	}
	utils.ReturnJson(ctx, suerror.SUCCESS, allCategory)
}
func GetSpecification(ctx *gin.Context) {

}
func GetShopcart(ctx *gin.Context) {
	type Date struct {
		Count       int64            `json:"count"`
		CurrentPage int              `json:"currentPage"`
		PageSize    int64              `json:"pageSize"`
		TotalPages  int64            `json:"totalPages"`
		Carts    *[]model.Cart `json:"data"`
	}
	var data Date

	type Param struct {
		Page int    `form:"page"`
		Name string `form:"name"`
	}
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		suParam.Page=1
		suParam.Name=""
		//utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		//return
	}
	var cart model.Cart

	data.Count,_=cart.GetShopCartCountNum()
	data.CurrentPage=suParam.Page
	data.PageSize=10
	data.TotalPages=(data.Count/data.PageSize)+1
	cartsWithPage, err := cart.GetShopCartsWithPage(suParam.Page)
	if err != nil {
		return
	}
	data.Carts=&cartsWithPage

	utils.ReturnJson(ctx,suerror.SUCCESS,data)

}
func GetMain(ctx *gin.Context) {
	type Param struct {
		Pindex int `form:"pindex"`
	}
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}

}
func GetGetQiniuToken(ctx *gin.Context) {

}
func GetFreight(ctx *gin.Context) {
	var freight model.FreightTemplate
	templates, err := freight.GetAllFreightTemplates()
	if err != nil {
		return
	}
	utils.ReturnJson(ctx, suerror.SUCCESS, templates)
}
func GetList(ctx *gin.Context) {

	type Date struct {
		Count       int64            `json:"count"`
		CurrentPage int              `json:"currentPage"`
		PageSize    int              `json:"pageSize"`
		TotalPages  int64            `json:"totalPages"`
		Shippers    *[]model.Shipper `json:"data"`
	}

	type Param struct {
		Page int    `form:"page"`
		Name string `form:"name"`
	}
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}
	var data Date
	var shipper model.Shipper

	if len(suParam.Name) > 0 {
		fmt.Println(suParam.Name)
		shipperCount, err := shipper.GetShipperCount()
		if err != nil {
			return
		}
		data.Count = shipperCount
		shippersWithPage, err2 := shipper.GetShippersWithPageAndName(suParam.Page, suParam.Name)
		if err2 != nil {
			return
		}
		data.CurrentPage = suParam.Page
		data.PageSize = 10
		data.Shippers = &shippersWithPage
		data.TotalPages = (shipperCount / 10) + 1
	} else {
		shipperCount, err := shipper.GetShipperCount()
		if err != nil {
			return
		}
		data.Count = shipperCount
		shippersWithPage, err2 := shipper.GetShippersWithPage(suParam.Page)
		if err2 != nil {
			return
		}
		data.CurrentPage = suParam.Page
		data.PageSize = 10
		data.Shippers = &shippersWithPage
		data.TotalPages = (shipperCount / 10) + 1

	}

	utils.ReturnJson(ctx, suerror.SUCCESS, data)

}
func GetGetareadata(ctx *gin.Context) {

}
func GetInfo(ctx *gin.Context) {
	type Param struct {
		Id int64 `form:"id"`
	}
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}
	var shipper model.Shipper
	shipperWithId, err := shipper.GetShipperWithId(suParam.Id)
	if err != nil {
		return
	}
	utils.ReturnJson(ctx, suerror.SUCCESS, shipperWithId)

}

//http://127.0.0.1:8360/admin/goods/onsale?page=1&size=10
func GetOnsale(ctx *gin.Context) {
	type Param struct {
		Page int `form:"page"`
		Size int `form:"size"`
	}
	type Date struct {
		Count       int            `json:"count"`
		CurrentPage int            `json:"currentPage"`
		PageSize    int            `json:"pageSize"`
		TotalPages  int            `json:"totalPages"`
		Goods       *[]model.Goods `json:"data"`
	}
	var data Date
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}
	var good model.Goods
	onSaleGoods, err := good.GetOnSaleGoods()
	if err != nil {
		return
	}
	data.Count = len(onSaleGoods)
	data.CurrentPage = suParam.Page
	data.PageSize = suParam.Size
	data.TotalPages = (data.Count / 10) + 1
	data.Goods = &onSaleGoods

	utils.ReturnJson(ctx, suerror.SUCCESS, data)

}
func GetOut(ctx *gin.Context) {
	type Param struct {
		Page int `form:"page"`
	}
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}
}
func GetDrop(ctx *gin.Context) {
	type Param struct {
		Page int `form:"page"`
	}
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}
}
func GetSort(ctx *gin.Context) {
	type Param struct {
		Page  int `form:"page"`
		Index int `form:"index"`
	}
	var suParam Param
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}
}

//func UserGet(ctx *gin.Context) {
//
//	type param struct {
//		ID int `form:"id" json:"id" binding:"required"`
//	}
//
//	var suParam param
//	if err := ctx.ShouldBindJSON(&suParam); err != nil {
//		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
//		return
//	}
//
//	userModel := model.SuUser{}
//
//	err := userModel.GetUserForID(suParam.ID)
//	if err != nil {
//		utils.ReturnJson(ctx, suerror.GORMSQLERROR, nil)
//		return
//	}
//
//	userModelJson, err := json.Marshal(&userModel)
//	if err != nil {
//		utils.ReturnJson(ctx, suerror.SUCCESS, userModelJson)
//		return
//	}
//}
//
//func Login(ctx *gin.Context) {
//
//	var suParam model.LoginParam
//	if err := ctx.ShouldBind(&suParam); err != nil {
//		utils.ReturnJson(ctx, suerror.PARAMNOTFUND, nil)
//		return
//	}
//	sign := utils.CheckSign(suParam)
//	if !sign {
//		utils.ReturnJson(ctx, suerror.SIGNERROR, nil)
//		return
//	}
//	fmt.Println(sign)
//	userModel := model.SuUser{}
//
//	err := userModel.GetUserForUserName(suParam.Username)
//	if err != nil {
//		utils.ReturnJson(ctx, suerror.SYSTEMERROR, nil)
//		return
//	}
//	//将前端传递的密码加盐计算MD5
//	userModel.UserPassword, _ = utils.SuMd5(userModel.UserPassword, configs.SuJwt.JwtSecret)
//	if userModel.UserName == suParam.Username && userModel.UserPassword == suParam.Password {
//		AccessToken, err := utils.CreateAccessToken(utils.WithUID(userModel.UserId))
//		if err != nil {
//			fmt.Println(err)
//		}
//		utils.ReturnJson(ctx, suerror.SUCCESS, []byte(AccessToken))
//		return
//	}
//	utils.ReturnJson(ctx, suerror.NOTFUND, nil)
//	return
//}
//
//func Register(ctx *gin.Context) {
//	var suparam model.RegisterParam
//	if err := ctx.ShouldBind(&suparam); err != nil {
//		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
//		return
//	}
//	sign := utils.CheckSign(suparam)
//	if !sign {
//		utils.ReturnJson(ctx, suerror.SIGNERROR, nil)
//		return
//	}
//	fmt.Println(sign)
//	//创建用户
//	userModel := model.SuUser{}
//
//	_, err := userModel.Create(suparam)
//	if err != nil {
//		utils.ReturnJson(ctx, suerror.GORMSQLERROR, nil)
//		return
//	}
//	utils.ReturnJson(ctx, suerror.SUCCESS, nil)
//	return
//}

func Test(ctx *gin.Context) {
	_ = utils.ErrorMail("fuck")
}
