package model

import "github.com/ilovesusu/su-gin/db"

type Settings struct {
	ID                 int64  `gorm:"column:id" json:"id" form:"id"`
	Autodelivery       int64  `gorm:"column:autodelivery" json:"autodelivery" form:"autodelivery"`
	Name               string `gorm:"column:name" json:"name" form:"name"`
	Tel                string `gorm:"column:tel" json:"tel" form:"tel"`
	Provincename       string `gorm:"column:provincename" json:"provincename" form:"provincename"`
	Cityname           string `gorm:"column:cityname" json:"cityname" form:"cityname"`
	Expareaname        string `gorm:"column:expareaname" json:"expareaname" form:"expareaname"`
	Address            string `gorm:"column:address" json:"address" form:"address"`
	DiscoveryImgHeight int64  `gorm:"column:discovery_img_height" json:"discovery_img_height" form:"discovery_img_height"`
	DiscoveryImg       string `gorm:"column:discovery_img" json:"discovery_img" form:"discovery_img"`
	GoodsId            int64  `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	CityId             int64  `gorm:"column:city_id" json:"city_id" form:"city_id"`
	ProvinceId         int64  `gorm:"column:province_id" json:"province_id" form:"province_id"`
	DistrictId         int64  `gorm:"column:district_id" json:"district_id" form:"district_id"`
	Countdown          int64  `gorm:"column:countdown" json:"countdown" form:"countdown"`
	Reset              int64  `gorm:"column:reset" json:"reset" form:"reset"`
}

func (model *Settings)GetOneSetting() (Setting Settings,err error)  {
	db.SuDB.Debug().Find(&Setting)
	return

}
