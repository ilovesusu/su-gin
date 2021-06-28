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
	ctx.ProtoBuf()
	type param struct {
		ID int `form:"id" json:"id" binding:"required"`
	}

	var suparam param
	if err := ctx.ShouldBindJSON(&suparam); err != nil {
		utils.ReturnJson(ctx, model.SuI18NList[40001].Error(), 40001, nil)
		return
	}

	userModel := model.User{}

	msg, err := userModel.GetUserForID(suparam.ID)
	if err != nil {
		utils.ReturnJson(ctx, msg, 40100, nil)
		return
	}

	userModeljson, err := json.Marshal(&userModel)
	if err != nil {
		utils.ReturnJson(ctx, model.SuI18NList[200].Error(), 200, userModeljson)
		return
	}
}

func Login(ctx *gin.Context) {

	var suparam model.LoginParam
	if err := ctx.ShouldBind(&suparam); err != nil {
		utils.ReturnJson(ctx, suerror.PARAMNOTFUND, nil)
		return
	}
	sign := utils.CheckSign(suparam)
	if !sign {
		utils.ReturnJson(ctx, suerror.SIGNERROR, nil)
		return
	}
	fmt.Println(sign)
	userModel := model.User{}

	msg, err := userModel.GetUserForUserName(suparam.Username)
	if err != nil {
		utils.ReturnJson(ctx, suerror.SYSTEMERROR, nil)
		return
	}
	//将前端传递的密码加盐计算MD5
	userModel.UserPassword, _ = utils.SuMd5(userModel.UserPassword, configs.SuJwt.JwtSecret)
	if userModel.UserName == suparam.Username && userModel.UserPassword == suparam.Password {
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
		utils.ReturnJson(ctx, model.SuI18NList[40001].Error()+err.Error(), 40001, nil)
		return
	}
	sign := utils.CheckSign(suparam)
	if !sign {
		utils.ReturnJson(ctx, model.SuI18NList[40004].Error(), 40004, nil)
		return
	}
	fmt.Println(sign)
	//创建用户
	userModel := model.User{}

	_, err := userModel.Create(suparam)
	if err != nil {
		utils.ReturnJson(ctx, model.SuI18NList[40100].Error()+err.Error(), 40100, nil)
		return
	}
	utils.ReturnJson(ctx, model.SuI18NList[200].Error()+model.SuI18NList[50001].Error(), 200, nil)
	return
}

func Test(ctx *gin.Context) {
	_ = utils.ErrorMail("fuck")
}
