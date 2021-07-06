package model

import (
	"github.com/ilovesusu/su-gin/db"
	"log"
)

//var Users []User
type User struct {
	ID            int64  `gorm:"column:id" json:"id" form:"id"`
	Nickname      string `gorm:"column:nickname" json:"nickname" form:"nickname"`
	Name          string `gorm:"column:name" json:"name" form:"name"`
	Username      string `gorm:"column:username" json:"username" form:"username"`
	Password      string `gorm:"column:password" json:"password" form:"password"`
	Gender        int64  `gorm:"column:gender" json:"gender" form:"gender"`
	Birthday      int64  `gorm:"column:birthday" json:"birthday" form:"birthday"`
	RegisterTime  int64  `gorm:"column:register_time" json:"register_time" form:"register_time"`
	LastLoginTime int64  `gorm:"column:last_login_time" json:"last_login_time" form:"last_login_time"`
	LastLoginIp   string `gorm:"column:last_login_ip" json:"last_login_ip" form:"last_login_ip"`
	Mobile        string `gorm:"column:mobile" json:"mobile" form:"mobile"`
	RegisterIp    string `gorm:"column:register_ip" json:"register_ip" form:"register_ip"`
	Avatar        string `gorm:"column:avatar" json:"avatar" form:"avatar"`
	WeixinOpenid  string `gorm:"column:weixin_openid" json:"weixin_openid" form:"weixin_openid"`
	NameMobile    int64  `gorm:"column:name_mobile" json:"name_mobile" form:"name_mobile"`
	Country       string `gorm:"column:country" json:"country" form:"country"`
	Province      string `gorm:"column:province" json:"province" form:"province"`
	City          string `gorm:"column:city" json:"city" form:"city"`
}

func (model *User) GetUserWithId(id int) (err error) {
	find := db.SuDB.Where("id=?", id).Find(&model)

	if find.Error != nil {
		log.Println(find.Error)
		err = find.Error
		return
	}

	return
}

func (model *User) GetUserWithPage(page int) (users []User, err error) {
	db.SuDB.Debug().Omit("password").Limit(10).Offset((page - 1) * 10).Find(&users)
	return
}
func (model *User)GetUserCountNum()  (num int64,err error){
	db.SuDB.Debug().Model(&User{}).Count(&num)
	return
}
