package model

type ExceptAreaDetail struct {
	ID           int64 `gorm:"column:id" json:"id" form:"id"`
	ExceptAreaId int64 `gorm:"column:except_area_id" json:"except_area_id" form:"except_area_id"`
	Area         int64 `gorm:"column:area" json:"area" form:"area"`
	IsDelete     int64 `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}