package model

import "github.com/ilovesusu/su-gin/db"

type Notice struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	Content  string `gorm:"column:content" json:"content" form:"content"`
	EndTime  int64  `gorm:"column:end_time" json:"end_time" form:"end_time"`
	IsDelete int64  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}
func (model *Notice)GetAllNotices() (notices []Notice,err error) {
	db.SuDB.Debug().Omit("password").Find(&notices)
	return
}