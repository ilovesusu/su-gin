package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilovesusu/su-gin/model"
	"github.com/ilovesusu/su-gin/suerror"
	"github.com/ilovesusu/su-gin/utils"
)

func GetGetAllRegion(ctx *gin.Context) {
	var regions model.Region
	all, err := regions.GetAll()
	if err != nil {
		return
	}
	utils.ReturnJson(ctx,suerror.SYSTEMERROR,all)

}
func GetGetAutoStatus(ctx *gin.Context) {
	utils.ReturnJson(ctx,suerror.SUCCESS,1)
}


