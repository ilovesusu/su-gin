package model

import "github.com/ilovesusu/su-gin/db"

type Admin struct {
	ID            int64  `gorm:"column:id" json:"id" form:"id"`
	Username      string `gorm:"column:username" json:"username" form:"username"`
	Password      string `gorm:"column:password" json:"password" form:"password"`
	PasswordSalt  string `gorm:"column:password_salt" json:"password_salt" form:"password_salt"`
	LastLoginIp   string `gorm:"column:last_login_ip" json:"last_login_ip" form:"last_login_ip"`
	LastLoginTime int64  `gorm:"column:last_login_time" json:"last_login_time" form:"last_login_time"`
	IsDelete      int64  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}

func (model *Admin)GetAllAdmins() (admins []Admin,err error) {
	db.SuDB.Debug().Omit("password").Find(&admins)
	return
}