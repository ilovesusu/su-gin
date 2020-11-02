package model

import (
	"github.com/ilovesusu/su-gin/config"
	"github.com/ilovesusu/su-gin/utils"
	"log"
)

type User struct {
	UserId       int    `form:"user_id" json:"user_id" primaryKey:"true"`
	UserName     string `form:"user_name" json:"user_name" gorm:"type:varchar(16);comment:'用户名'"`
	UserPassword string `form:"user_password" json:"user_password" gorm:"type:varchar(16);comment:'密码'"`
}

type LoginParam struct {
	Username  string `form:"username" susu:"username" json:"username" binding:"required"`
	Password  string `form:"password" susu:"password" json:"password" binding:"required"`
	SecretKey string `form:"secret_key" susu:"-" json:"secret_key" binding:"required"`
	TimeStamp uint   `form:"time_stamp" susu:"time_stamp" json:"time_stamp" binding:"required"`
}

type RegisterParam struct {
	Username  string `form:"username" susu:"username" json:"username" binding:"required"`
	Password  string `form:"password" susu:"password" json:"password" binding:"required"`
	SecretKey string `form:"secret_key" susu:"-" json:"secret_key" binding:"required"`
	TimeStamp uint   `form:"time_stamp" susu:"time_stamp" json:"time_stamp" binding:"required"`
}

func (model *User) GetUserForID(id int) (user User, err error) {
	find := SuDB.Where("id=?", id).Find(&user)

	if find.Error != nil {
		log.Println(find.Error)
		err = find.Error
		return
	}

	return
}

func (model *User) GetUserForUserName(username string) (user User, err error) {
	find := SuDB.Where("user_name=?", username).Find(&user)

	if find.Error != nil {
		log.Println(find.Error)
		err = find.Error
		return
	}

	return
}

func (model *User) CreateUser(registerparam RegisterParam) (user User, err error) {
	user.UserPassword, _ = utils.SuMd5(registerparam.Password, config.SuApp.JwtSecret)
	user.UserName = registerparam.Username

	find := SuDB.Create(&user)

	if find.Error != nil {
		log.Println(find.Error)
		err = find.Error
		return
	}
	return
}
