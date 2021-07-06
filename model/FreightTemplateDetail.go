package model

type FreightTemplateDetail struct {
	ID         int64 `gorm:"column:id" json:"id" form:"id"`
	TemplateId int64 `gorm:"column:template_id" json:"template_id" form:"template_id"`
	GroupId    int64 `gorm:"column:group_id" json:"group_id" form:"group_id"`
	Area       int64 `gorm:"column:area" json:"area" form:"area"`
	IsDelete   int64 `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}
