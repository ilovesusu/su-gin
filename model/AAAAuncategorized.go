package model

type gfdagAd struct {
	UserId       int    `form:"user_id" json:"user_id" primaryKey:"true"`
	UserName     string `form:"user_name" json:"user_name" gorm:"type:varchar(16);comment:'用户名'"`
	UserPassword string `form:"user_password" json:"user_password" gorm:"type:varchar(16);comment:'密码'"`

	Id uint `form:"log_id" json:"log_id" gorm:"primary_key"`
}



















