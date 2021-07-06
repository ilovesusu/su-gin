package model
type Specification struct {
	ID        int64  `gorm:"column:id" json:"id" form:"id"`
	Name      string `gorm:"column:name" json:"name" form:"name"`
	SortOrder int64  `gorm:"column:sort_order" json:"sort_order" form:"sort_order"`
	Memo      string `gorm:"column:memo" json:"memo" form:"memo"`
}

