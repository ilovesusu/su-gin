package model
type SearchHistory struct {
	ID      int64  `gorm:"column:id" json:"id" form:"id"`
	Keyword string `gorm:"column:keyword" json:"keyword" form:"keyword"`
	From    string `gorm:"column:from" json:"from" form:"from"`
	AddTime int64  `gorm:"column:add_time" json:"add_time" form:"add_time"`
	UserId  string `gorm:"column:user_id" json:"user_id" form:"user_id"`
}

