package model
type Footprint struct {
	ID      int64 `gorm:"column:id" json:"id" form:"id"`
	UserId  int64 `gorm:"column:user_id" json:"user_id" form:"user_id"`
	GoodsId int64 `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	AddTime int64 `gorm:"column:add_time" json:"add_time" form:"add_time"`
}