package model
type Keywords struct {
	Keyword   string `gorm:"column:keyword" json:"keyword" form:"keyword"`
	IsHot     int64  `gorm:"column:is_hot" json:"is_hot" form:"is_hot"`
	IsDefault int64  `gorm:"column:is_default" json:"is_default" form:"is_default"`
	IsShow    int64  `gorm:"column:is_show" json:"is_show" form:"is_show"`
	SortOrder int64  `gorm:"column:sort_order" json:"sort_order" form:"sort_order"`
	ID        int64  `gorm:"column:id" json:"id" form:"id"`
	Type      int64  `gorm:"column:type" json:"type" form:"type"`
}
