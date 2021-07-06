package model
type FreightTemplateGroup struct {
	ID           int64   `gorm:"column:id" json:"id" form:"id"`
	TemplateId   int64   `gorm:"column:template_id" json:"template_id" form:"template_id"`
	IsDefault    int64   `gorm:"column:is_default" json:"is_default" form:"is_default"`
	Area         string  `gorm:"column:area" json:"area" form:"area"`
	Start        int64   `gorm:"column:start" json:"start" form:"start"`
	StartFee     float64 `gorm:"column:start_fee" json:"start_fee" form:"start_fee"`
	Add          int64   `gorm:"column:add" json:"add" form:"add"`
	AddFee       float64 `gorm:"column:add_fee" json:"add_fee" form:"add_fee"`
	FreeByNumber int64   `gorm:"column:free_by_number" json:"free_by_number" form:"free_by_number"`
	FreeByMoney  float64 `gorm:"column:free_by_money" json:"free_by_money" form:"free_by_money"`
	IsDelete     int64   `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}
