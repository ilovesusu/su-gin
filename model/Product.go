package model
type Product struct {
	ID                    int64   `gorm:"column:id" json:"id" form:"id"`
	GoodsId               int64   `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	GoodsSpecificationIds string  `gorm:"column:goods_specification_ids" json:"goods_specification_ids" form:"goods_specification_ids"`
	GoodsSn               string  `gorm:"column:goods_sn" json:"goods_sn" form:"goods_sn"`
	GoodsNumber           int64   `gorm:"column:goods_number" json:"goods_number" form:"goods_number"`
	RetailPrice           float64 `gorm:"column:retail_price" json:"retail_price" form:"retail_price"`
	Cost                  float64 `gorm:"column:cost" json:"cost" form:"cost"`
	GoodsWeight           float64 `gorm:"column:goods_weight" json:"goods_weight" form:"goods_weight"`
	HasChange             int64   `gorm:"column:has_change" json:"has_change" form:"has_change"`
	GoodsName             string  `gorm:"column:goods_name" json:"goods_name" form:"goods_name"`
	IsOnSale              int64   `gorm:"column:is_on_sale" json:"is_on_sale" form:"is_on_sale"`
	IsDelete              int64   `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}

