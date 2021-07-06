package model
type GoodsGallery struct {
	ID        int64  `gorm:"column:id" json:"id" form:"id"`
	GoodsId   int64  `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	ImgUrl    string `gorm:"column:img_url" json:"img_url" form:"img_url"`
	ImgDesc   string `gorm:"column:img_desc" json:"img_desc" form:"img_desc"`
	SortOrder int64  `gorm:"column:sort_order" json:"sort_order" form:"sort_order"`
	IsDelete  int64  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}
