package model

import (
	"github.com/ilovesusu/su-gin/configs"
	"github.com/ilovesusu/su-gin/db"
	"github.com/ilovesusu/su-gin/utils"
	"log"
)

type User struct {
	UserId       int    `form:"user_id" json:"user_id" primaryKey:"true"`
	UserName     string `form:"user_name" json:"user_name" gorm:"type:varchar(16);comment:'用户名'"`
	UserPassword string `form:"user_password" json:"user_password" gorm:"type:varchar(16);comment:'密码'"`
}

type LoginParam struct {
	Username  string `form:"username" example:"susu" susu:"username" json:"username" binding:"required"`
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

func init() {
	//自动迁移表,AutoMigrate 会创建表，缺少的外键，约束，列和索引，并且会更改现有列的类型（如果其大小、精度、是否为空可更改）。但 不会 删除未使用的列，以保护您的数据。
	_ = db.SuDB.AutoMigrate(&User{}) //用户表
}

func (user *User) Create(registerParam RegisterParam) (msg string, err error) {
	user.UserPassword, _ = utils.SuMd5(registerParam.Password, configs.SuJwt.JwtSecret)
	user.UserName = registerParam.Username

	find := db.SuDB.Create(&user)

	if find.Error != nil {
		log.Println(find.Error)
		err = find.Error
		return
	}
	return
}

func (user *User) Retrieve() (msg string, err error) {
	panic("implement me")
}

func (user *User) Update() (msg string, err error) {
	panic("implement me")
}

func (user *User) Delect() (msg string, err error) {
	panic("implement me")
}

func (model *User) GetUserForID(id int) (err error) {
	find := db.SuDB.Where("id=?", id).Find(&model)

	if find.Error != nil {
		log.Println(find.Error)
		err = find.Error
		return
	}

	return
}

func (model *User) GetUserForUserName(username string) (err error) {
	find := db.SuDB.Where("user_name=?", username).Find(&model)

	if find.Error != nil {
		log.Println(find.Error)
		err = find.Error
		return
	}

	return
}
