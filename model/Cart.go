package model

import "github.com/ilovesusu/su-gin/db"

type Cart struct {
	ID                        int64   `gorm:"column:id" json:"id" form:"id"`
	UserId                    int64   `gorm:"column:user_id" json:"user_id" form:"user_id"`
	GoodsId                   int64   `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	GoodsSn                   string  `gorm:"column:goods_sn" json:"goods_sn" form:"goods_sn"`
	ProductId                 int64   `gorm:"column:product_id" json:"product_id" form:"product_id"`
	GoodsName                 string  `gorm:"column:goods_name" json:"goods_name" form:"goods_name"`
	GoodsAka                  string  `gorm:"column:goods_aka" json:"goods_aka" form:"goods_aka"`
	GoodsWeight               float64 `gorm:"column:goods_weight" json:"goods_weight" form:"goods_weight"`
	AddPrice                  float64 `gorm:"column:add_price" json:"add_price" form:"add_price"`
	RetailPrice               float64 `gorm:"column:retail_price" json:"retail_price" form:"retail_price"`
	Number                    int64   `gorm:"column:number" json:"number" form:"number"`
	GoodsSpecifitionNameValue string  `gorm:"column:goods_specifition_name_value" json:"goods_specifition_name_value" form:"goods_specifition_name_value"`
	GoodsSpecifitionIds       string  `gorm:"column:goods_specifition_ids" json:"goods_specifition_ids" form:"goods_specifition_ids"`
	Checked                   int64   `gorm:"column:checked" json:"checked" form:"checked"`
	ListPicUrl                string  `gorm:"column:list_pic_url" json:"list_pic_url" form:"list_pic_url"`
	FreightTemplateId         int64   `gorm:"column:freight_template_id" json:"freight_template_id" form:"freight_template_id"`
	IsOnSale                  int64   `gorm:"column:is_on_sale" json:"is_on_sale" form:"is_on_sale"`
	AddTime                   int64   `gorm:"column:add_time" json:"add_time" form:"add_time"`
	IsFast                    int64   `gorm:"column:is_fast" json:"is_fast" form:"is_fast"`
	IsDelete                  int64   `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}

func (model *Cart) GetShopCartCountNum() (num int64, err error) {
	db.SuDB.Debug().Model(&Cart{}).Count(&num)
	return
}
func (model *Cart) GetShopCartsWithPage(page int) (shopCartsWithPage []Cart, err error) {
	db.SuDB.Debug().Model(&Cart{}).Limit(10).Offset((page - 1) * 10).Order("add_time").Find(&shopCartsWithPage)
	return
}
