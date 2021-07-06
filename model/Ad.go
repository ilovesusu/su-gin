package model

import "github.com/ilovesusu/su-gin/db"

type Ad struct {
	ID        int64  `gorm:"column:id" json:"id" form:"id"`
	LinkType  int64  `gorm:"column:link_type" json:"link_type" form:"link_type"`
	Link      string `gorm:"column:link" json:"link" form:"link"`
	GoodsId   int64  `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	ImageUrl  string `gorm:"column:image_url" json:"image_url" form:"image_url"`
	EndTime   int64  `gorm:"column:end_time" json:"end_time" form:"end_time"`
	Enabled   int64  `gorm:"column:enabled" json:"enabled" form:"enabled"`
	SortOrder int64  `gorm:"column:sort_order" json:"sort_order" form:"sort_order"`
	IsDelete  int64  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}

func (model *Ad)GetAdsWithPage(page int) (ads []Ad,err error) {
	db.SuDB.Debug().Omit("password").Limit(10).Offset((page - 1) * 10).Find(&ads)
	return
}
