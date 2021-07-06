package model
type GoodsSpecification struct {
	ID              int64  `gorm:"column:id" json:"id" form:"id"`
	GoodsId         int64  `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	SpecificationId int64  `gorm:"column:specification_id" json:"specification_id" form:"specification_id"`
	Value           string `gorm:"column:value" json:"value" form:"value"`
	PicUrl          string `gorm:"column:pic_url" json:"pic_url" form:"pic_url"`
	IsDelete        int64  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}
