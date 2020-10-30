package model

import (
	"database/sql"
	"fmt"
	"github.com/ilovesusu/su-gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var (
	db   *sql.DB
	SuDB *gorm.DB
	err  error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.SuDatabase.User,
		config.SuDatabase.PassWord,
		config.SuDatabase.Host,
		config.SuDatabase.Port,
		config.SuDatabase.DatabaseName)
	SuDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.SuDatabase.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,                          // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//SuDB, err = gorm.Open(mysql.Open("root:password0@tcp(xx.xx.xx.xx)/lr_table?charset=utf8&parseTime=True&loc=Local"),&gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	db, _ := SuDB.DB()

	//连接池
	//通过 show processlist; 查看连接数
	db.SetMaxIdleConns(config.SuDatabase.MAX_IDLE_CONNS)
	db.SetMaxOpenConns(config.SuDatabase.MAX_OPEN_CONNS)
	db.SetConnMaxLifetime(time.Minute * config.SuDatabase.CONN_MAX_LIFE_TIME)

	//自动迁移表,AutoMigrate 会创建表，缺少的外键，约束，列和索引，并且会更改现有列的类型（如果其大小、精度、是否为空可更改）。但 不会 删除未使用的列，以保护您的数据。
	_ = SuDB.AutoMigrate(&User{}) //用户表
}
