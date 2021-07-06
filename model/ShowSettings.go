package model

import "github.com/ilovesusu/su-gin/db"

type ShowSettings struct {
	ID             int64 `gorm:"column:id" json:"id" form:"id"`
	Banner         int64 `gorm:"column:banner" json:"banner" form:"banner"`
	Channel        int64 `gorm:"column:channel" json:"channel" form:"channel"`
	IndexBannerImg int64 `gorm:"column:index_banner_img" json:"index_banner_img" form:"index_banner_img"`
	Notice         int64 `gorm:"column:notice" json:"notice" form:"notice"`
}

func (s *ShowSettings)GetShowSettings() (err error) {
	db.SuDB.Debug().Model(&ShowSettings{}).Find(&s)
	return
}
