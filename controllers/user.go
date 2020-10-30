package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/config"
	"github.com/ilovesusu/su-gin/model"
	"github.com/ilovesusu/su-gin/utils"
)

func UserGet(ctx *gin.Context) {
	type param struct {
		ID int `form:"id" json:"id" binding:"required"`
	}

	var suparam param
	if err := ctx.ShouldBindJSON(&suparam); err != nil {
		utils.ReturnJson(ctx, model.SuI18NList[40001].Error(), 40001, nil)
		return
	}

	userModel := model.User{}

	user, err := userModel.GetUserForID(suparam.ID)
	if err != nil {
		utils.ReturnJson(ctx, model.SuI18NList[40100].Error()+err.Error(), 40100, nil)
		return
	}

	json_str, err := json.Marshal(&user)
	if err != nil {
		utils.ReturnJson(ctx, model.SuI18NList[200].Error(), 200, json_str)
		return
	}
}

func Login(ctx *gin.Context) {

	var suparam model.LoginParam
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
	userModel := model.User{}

	user, err := userModel.GetUserForUserName(suparam.Username)
	if err != nil {
		utils.ReturnJson(ctx, model.SuI18NList[40100].Error()+err.Error(), 40100, nil)
		return
	}
	//将前端传递的密码加盐计算MD5
	user.UserPassword, _ = utils.SuMd5(user.UserPassword, config.SuApp.JwtSecret)
	if user.UserName == suparam.Username && user.UserPassword == suparam.Password {
		utils.ReturnJson(ctx, model.SuI18NList[200].Error(), 200, nil)
		return
	}
	utils.ReturnJson(ctx, model.SuI18NList[40000].Error(), 40000, nil)
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

	_, err := userModel.CreateUser(suparam)
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