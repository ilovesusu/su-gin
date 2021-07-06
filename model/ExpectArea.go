package model
type ExceptArea struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	Content  string `gorm:"column:content" json:"content" form:"content"`
	Area     string `gorm:"column:area" json:"area" form:"area"`
	IsDelete int64  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}
