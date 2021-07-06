package model
type Formid struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	UserId   int64  `gorm:"column:user_id" json:"user_id" form:"user_id"`
	OrderId  int64  `gorm:"column:order_id" json:"order_id" form:"order_id"`
	FormId   string `gorm:"column:form_id" json:"form_id" form:"form_id"`
	AddTime  int64  `gorm:"column:add_time" json:"add_time" form:"add_time"`
	UseTimes int64  `gorm:"column:use_times" json:"use_times" form:"use_times"`
}