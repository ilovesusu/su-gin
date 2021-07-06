package model

import "github.com/ilovesusu/su-gin/db"

type Goods struct {
	ID                int64   `gorm:"column:id" json:"id" form:"id"`
	CategoryId        int64   `gorm:"column:category_id" json:"category_id" form:"category_id"`
	IsOnSale          int64   `gorm:"column:is_on_sale" json:"is_on_sale" form:"is_on_sale"`
	Name              string  `gorm:"column:name" json:"name" form:"name"`
	GoodsNumber       int64   `gorm:"column:goods_number" json:"goods_number" form:"goods_number"`
	SellVolume        int64   `gorm:"column:sell_volume" json:"sell_volume" form:"sell_volume"`
	Keywords          string  `gorm:"column:keywords" json:"keywords" form:"keywords"`
	RetailPrice       string  `gorm:"column:retail_price" json:"retail_price" form:"retail_price"`
	MinRetailPrice    float64 `gorm:"column:min_retail_price" json:"min_retail_price" form:"min_retail_price"`
	CostPrice         string  `gorm:"column:cost_price" json:"cost_price" form:"cost_price"`
	MinCostPrice      float64 `gorm:"column:min_cost_price" json:"min_cost_price" form:"min_cost_price"`
	GoodsBrief        string  `gorm:"column:goods_brief" json:"goods_brief" form:"goods_brief"`
	GoodsDesc         string  `gorm:"column:goods_desc" json:"goods_desc" form:"goods_desc"`
	SortOrder         int64   `gorm:"column:sort_order" json:"sort_order" form:"sort_order"`
	IsIndex           int64   `gorm:"column:is_index" json:"is_index" form:"is_index"`
	IsNew             int64   `gorm:"column:is_new" json:"is_new" form:"is_new"`
	GoodsUnit         string  `gorm:"column:goods_unit" json:"goods_unit" form:"goods_unit"`
	HttpsPicUrl       string  `gorm:"column:https_pic_url" json:"https_pic_url" form:"https_pic_url"`
	ListPicUrl        string  `gorm:"column:list_pic_url" json:"list_pic_url" form:"list_pic_url"`
	FreightTemplateId int64   `gorm:"column:freight_template_id" json:"freight_template_id" form:"freight_template_id"`
	FreightType       int64   `gorm:"column:freight_type" json:"freight_type" form:"freight_type"`
	IsDelete          int64   `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}
func (model *Goods) GetGoodsOnsaleNumber() (num int64,err error) {
	//接收器入，gorm使用的是Goods,与接收器类型失去了联系,(已改进
	db.SuDB.Debug().Model(&Goods{}).Where("is_on_sale",1).Count(&num)
	return
}
func (model *Goods) GetOnSaleGoods() (goodsOnSale []Goods,err error) {
	db.SuDB.Debug().Model(&Goods{}).Where("is_on_sale",1).Find(&goodsOnSale)
	return
}
func (model *Goods)GetAllGoods() (goodsOnSale []Goods,err error) {
	db.SuDB.Debug().Model(&Goods{}).Find(&goodsOnSale)
	return
}