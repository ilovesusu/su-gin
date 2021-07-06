package model

import "github.com/ilovesusu/su-gin/db"

type FreightTemplate struct {
	ID           int64   `gorm:"column:id" json:"id" form:"id"`
	Name         string  `gorm:"column:name" json:"name" form:"name"`
	PackagePrice float64 `gorm:"column:package_price" json:"package_price" form:"package_price"`
	FreightType  int64   `gorm:"column:freight_type" json:"freight_type" form:"freight_type"`
	IsDelete     int64   `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}

func (model *FreightTemplate) GetAllFreightTemplates() (freightTemplates []FreightTemplate,err error)  {
	db.SuDB.Debug().Find(&freightTemplates)
	return
}
