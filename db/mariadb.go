package db

import (
	"fmt"
	"github.com/ilovesusu/su-gin/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var (
	SuDB *gorm.DB
	err  error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.SuDatabase.User,
		configs.SuDatabase.PassWord,
		configs.SuDatabase.Host,
		configs.SuDatabase.Port,
		configs.SuDatabase.DatabaseName)
	SuDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   configs.SuDatabase.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,                           // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	db, _ := SuDB.DB()

	//连接池
	//通过 show processlist; 查看连接数
	db.SetMaxIdleConns(configs.SuDatabase.MAX_IDLE_CONNS)
	db.SetMaxOpenConns(configs.SuDatabase.MAX_OPEN_CONNS)
	db.SetConnMaxLifetime(time.Minute * configs.SuDatabase.CONN_MAX_LIFE_TIME)
}
