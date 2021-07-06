package model
type OrderGoods struct {
	ID                        int64   `gorm:"column:id" json:"id" form:"id"`
	OrderId                   int64   `gorm:"column:order_id" json:"order_id" form:"order_id"`
	GoodsId                   int64   `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	GoodsName                 string  `gorm:"column:goods_name" json:"goods_name" form:"goods_name"`
	GoodsAka                  string  `gorm:"column:goods_aka" json:"goods_aka" form:"goods_aka"`
	ProductId                 int64   `gorm:"column:product_id" json:"product_id" form:"product_id"`
	Number                    int64   `gorm:"column:number" json:"number" form:"number"`
	RetailPrice               float64 `gorm:"column:retail_price" json:"retail_price" form:"retail_price"`
	GoodsSpecifitionNameValue string  `gorm:"column:goods_specifition_name_value" json:"goods_specifition_name_value" form:"goods_specifition_name_value"`
	GoodsSpecifitionIds       string  `gorm:"column:goods_specifition_ids" json:"goods_specifition_ids" form:"goods_specifition_ids"`
	ListPicUrl                string  `gorm:"column:list_pic_url" json:"list_pic_url" form:"list_pic_url"`
	UserId                    int64   `gorm:"column:user_id" json:"user_id" form:"user_id"`
	IsDelete                  int64   `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}

