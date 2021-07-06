package model

import "github.com/ilovesusu/su-gin/db"

type Region struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	ParentId int64  `gorm:"column:parent_id" json:"parent_id" form:"parent_id"`
	Name     string `gorm:"column:name" json:"name" form:"name"`
	Type     int64  `gorm:"column:type" json:"type" form:"type"`
	AgencyId int64  `gorm:"column:agency_id" json:"agency_id" form:"agency_id"`
	Area     int64  `gorm:"column:area" json:"area" form:"area"`
	AreaCode string `gorm:"column:area_code" json:"area_code" form:"area_code"`
	FarArea  int64  `gorm:"column:far_area" json:"far_area" form:"far_area"`
}

func (model *Region) GetAll() (regions []Region, err error) {
	db.SuDB.Debug().Find(&regions)
	return
}
