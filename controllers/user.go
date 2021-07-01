package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/configs"
	"github.com/ilovesusu/su-gin/model"
	"github.com/ilovesusu/su-gin/suerror"
	"github.com/ilovesusu/su-gin/utils"
)

func UserGet(ctx *gin.Context) {

	type param struct {
		ID int `form:"id" json:"id" binding:"required"`
	}

	var suParam param
	if err := ctx.ShouldBindJSON(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}

	userModel := model.User{}

	err := userModel.GetUserForID(suParam.ID)
	if err != nil {
		utils.ReturnJson(ctx, suerror.GORMSQLERROR, nil)
		return
	}

	userModelJson, err := json.Marshal(&userModel)
	if err != nil {
		utils.ReturnJson(ctx, suerror.SUCCESS, userModelJson)
		return
	}
}

func Login(ctx *gin.Context) {

	var suParam model.LoginParam
	if err := ctx.ShouldBind(&suParam); err != nil {
		utils.ReturnJson(ctx, suerror.PARAMNOTFUND, nil)
		return
	}
	sign := utils.CheckSign(suParam)
	if !sign {
		utils.ReturnJson(ctx, suerror.SIGNERROR, nil)
		return
	}
	fmt.Println(sign)
	userModel := model.User{}

	err := userModel.GetUserForUserName(suParam.Username)
	if err != nil {
		utils.ReturnJson(ctx, suerror.SYSTEMERROR, nil)
		return
	}
	//将前端传递的密码加盐计算MD5
	userModel.UserPassword, _ = utils.SuMd5(userModel.UserPassword, configs.SuJwt.JwtSecret)
	if userModel.UserName == suParam.Username && userModel.UserPassword == suParam.Password {
		AccessToken, err := utils.CreateAccessToken(utils.WithUID(userModel.UserId))
		if err != nil {
			fmt.Println(err)
		}
		utils.ReturnJson(ctx, suerror.SUCCESS, []byte(AccessToken))
		return
	}
	utils.ReturnJson(ctx, suerror.NOTFUND, nil)
	return
}

func Register(ctx *gin.Context) {
	var suparam model.RegisterParam
	if err := ctx.ShouldBind(&suparam); err != nil {
		utils.ReturnJson(ctx, suerror.NOPARAMETER, nil)
		return
	}
	sign := utils.CheckSign(suparam)
	if !sign {
		utils.ReturnJson(ctx, suerror.SIGNERROR, nil)
		return
	}
	fmt.Println(sign)
	//创建用户
	userModel := model.User{}

	_, err := userModel.Create(suparam)
	if err != nil {
		utils.ReturnJson(ctx, suerror.GORMSQLERROR, nil)
		return
	}
	utils.ReturnJson(ctx, suerror.SUCCESS, nil)
	return
}

func Test(ctx *gin.Context) {
	_ = utils.ErrorMail("fuck")
}
