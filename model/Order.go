package model

import "github.com/ilovesusu/su-gin/db"

type Order struct {
	ID             int64   `gorm:"column:id" json:"id" form:"id"`
	OrderSn        string  `gorm:"column:order_sn" json:"order_sn" form:"order_sn"`
	UserId         int64   `gorm:"column:user_id" json:"user_id" form:"user_id"`
	OrderStatus    int64   `gorm:"column:order_status" json:"order_status" form:"order_status"`
	OfflinePay     int64   `gorm:"column:offline_pay" json:"offline_pay" form:"offline_pay"`
	ShippingStatus int64   `gorm:"column:shipping_status" json:"shipping_status" form:"shipping_status"`
	PrintStatus    int64   `gorm:"column:print_status" json:"print_status" form:"print_status"`
	PayStatus      int64   `gorm:"column:pay_status" json:"pay_status" form:"pay_status"`
	Consignee      string  `gorm:"column:consignee" json:"consignee" form:"consignee"`
	Country        int64   `gorm:"column:country" json:"country" form:"country"`
	Province       int64   `gorm:"column:province" json:"province" form:"province"`
	City           int64   `gorm:"column:city" json:"city" form:"city"`
	District       int64   `gorm:"column:district" json:"district" form:"district"`
	Address        string  `gorm:"column:address" json:"address" form:"address"`
	PrintInfo      string  `gorm:"column:print_info" json:"print_info" form:"print_info"`
	Mobile         string  `gorm:"column:mobile" json:"mobile" form:"mobile"`
	Postscript     string  `gorm:"column:postscript" json:"postscript" form:"postscript"`
	AdminMemo      string  `gorm:"column:admin_memo" json:"admin_memo" form:"admin_memo"`
	ShippingFee    float64 `gorm:"column:shipping_fee" json:"shipping_fee" form:"shipping_fee"`
	PayName        string  `gorm:"column:pay_name" json:"pay_name" form:"pay_name"`
	PayId          string  `gorm:"column:pay_id" json:"pay_id" form:"pay_id"`
	ChangePrice    float64 `gorm:"column:change_price" json:"change_price" form:"change_price"`
	ActualPrice    float64 `gorm:"column:actual_price" json:"actual_price" form:"actual_price"`
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price" form:"order_price"`
	GoodsPrice     float64 `gorm:"column:goods_price" json:"goods_price" form:"goods_price"`
	AddTime        int64   `gorm:"column:add_time" json:"add_time" form:"add_time"`
	PayTime        int64   `gorm:"column:pay_time" json:"pay_time" form:"pay_time"`
	ShippingTime   int64   `gorm:"column:shipping_time" json:"shipping_time" form:"shipping_time"`
	ConfirmTime    int64   `gorm:"column:confirm_time" json:"confirm_time" form:"confirm_time"`
	DealdoneTime   int64   `gorm:"column:dealdone_time" json:"dealdone_time" form:"dealdone_time"`
	FreightPrice   int64   `gorm:"column:freight_price" json:"freight_price" form:"freight_price"`
	ExpressValue   float64 `gorm:"column:express_value" json:"express_value" form:"express_value"`
	Remark         string  `gorm:"column:remark" json:"remark" form:"remark"`
	OrderType      int64   `gorm:"column:order_type" json:"order_type" form:"order_type"`
	IsDelete       int64   `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}



func (model *User) GetqasdfghUrsturtserWithPage(page int) (users []User, err error) {
	//*user 作为接收器，返回参数是[]User?
	db.SuDB.Debug().Omit("password").Limit(10).Offset((page - 1) * 10).Find(&users)
	return
}

