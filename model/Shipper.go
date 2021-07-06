package model

import "github.com/ilovesusu/su-gin/db"

type Shipper struct {
	ID           int64  `gorm:"column:id" json:"id" form:"id"`
	Name         string `gorm:"column:name" json:"name" form:"name"`
	Code         string `gorm:"column:code" json:"code" form:"code"`
	SortOrder    int64  `gorm:"column:sort_order" json:"sort_order" form:"sort_order"`
	Monthcode    string `gorm:"column:monthcode" json:"monthcode" form:"monthcode"`
	Customername string `gorm:"column:customername" json:"customername" form:"customername"`
	Enabled      int64  `gorm:"column:enabled" json:"enabled" form:"enabled"`
}

func (model *Shipper) GetAllEnableShippers()(enableShipper []Shipper,err error)  {
	db.SuDB.Debug().Where("enabled=?",1).Find(&enableShipper)
	return
}
func (model *Shipper) GetShipperCount() (count int64,err error){
	db.SuDB.Debug().Model(&Shipper{}).Count(&count)
	return
}
func (model *Shipper) GetShippersWithPage(page int)(ShippersWithPage []Shipper,err error)  {
	db.SuDB.Debug().Limit(10).Offset((page - 1) * 10).Find(&ShippersWithPage)

	return
}
func (model *Shipper) GetShippersWithPageAndName(page int,name string)(ShippersWithPage []Shipper,err error)  {
	db.SuDB.Debug().Where("name LIKE ?", "%"+name+"%").Limit(10).Offset((page - 1) * 10).Find(&ShippersWithPage)
	return
}
func (model *Shipper)GetShipperWithId(Id int64) (ShipperWithId Shipper,err error) {
	db.SuDB.Debug().Where("id = ?",Id).Find(&ShipperWithId)
	return
}