package model
type OrderExpress struct {
	ID           int64  `gorm:"column:id" json:"id" form:"id"`
	OrderId      int64  `gorm:"column:order_id" json:"order_id" form:"order_id"`
	ShipperId    int64  `gorm:"column:shipper_id" json:"shipper_id" form:"shipper_id"`
	ShipperName  string `gorm:"column:shipper_name" json:"shipper_name" form:"shipper_name"`
	ShipperCode  string `gorm:"column:shipper_code" json:"shipper_code" form:"shipper_code"`
	LogisticCode string `gorm:"column:logistic_code" json:"logistic_code" form:"logistic_code"`
	Traces       string `gorm:"column:traces" json:"traces" form:"traces"`
	IsFinish     int64  `gorm:"column:is_finish" json:"is_finish" form:"is_finish"`
	RequestCount int64  `gorm:"column:request_count" json:"request_count" form:"request_count"`
	RequestTime  int64  `gorm:"column:request_time" json:"request_time" form:"request_time"`
	AddTime      int64  `gorm:"column:add_time" json:"add_time" form:"add_time"`
	UpdateTime   int64  `gorm:"column:update_time" json:"update_time" form:"update_time"`
	ExpressType  int64  `gorm:"column:express_type" json:"express_type" form:"express_type"`
	RegionCode   string `gorm:"column:region_code" json:"region_code" form:"region_code"`
}
