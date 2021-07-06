package model

import "github.com/ilovesusu/su-gin/db"

type Category struct {
	ID         int64  `gorm:"column:id" json:"id" form:"id"`
	Name       string `gorm:"column:name" json:"name" form:"name"`
	Keywords   string `gorm:"column:keywords" json:"keywords" form:"keywords"`
	FrontDesc  string `gorm:"column:front_desc" json:"front_desc" form:"front_desc"`
	ParentId   int64  `gorm:"column:parent_id" json:"parent_id" form:"parent_id"`
	SortOrder  int64  `gorm:"column:sort_order" json:"sort_order" form:"sort_order"`
	ShowIndex  int64  `gorm:"column:show_index" json:"show_index" form:"show_index"`
	IsShow     int64  `gorm:"column:is_show" json:"is_show" form:"is_show"`
	IconUrl    string `gorm:"column:icon_url" json:"icon_url" form:"icon_url"`
	ImgUrl     string `gorm:"column:img_url" json:"img_url" form:"img_url"`
	Level      string `gorm:"column:level" json:"level" form:"level"`
	FrontName  string `gorm:"column:front_name" json:"front_name" form:"front_name"`
	PHeight    int64  `gorm:"column:p_height" json:"p_height" form:"p_height"`
	IsCategory int64  `gorm:"column:is_category" json:"is_category" form:"is_category"`
	IsChannel  int64  `gorm:"column:is_channel" json:"is_channel" form:"is_channel"`
}

func (model *Category)GetAllCategory()(categorys []Category,err error)  {
	db.SuDB.Debug().Omit("password").Find(&categorys)
	return
}
func (model *Category)GetTopCategory()(categorys []Category,err error)  {
	db.SuDB.Debug().Limit(10).Find(&categorys)
	return
}