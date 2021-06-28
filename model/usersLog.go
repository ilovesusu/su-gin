package model

import (
	"github.com/ilovesusu/su-gin/db"
	"time"
)

//商家操作日志
type MerchantLog struct {
	LogId           uint      `form:"log_id" json:"log_id" gorm:"primary_key"`
	MerchantId      uint      `form:"merchant_id" json:"merchant_id" gorm:"comment:'操作人id'"`
	MerchantName    string    `form:"merchant_name" json:"merchant_name" gorm:"comment:'操作账号'"`
	OperatingText   string    `form:"operating_text" json:"operating_text" gorm:"type:text;comment:'操作内容'"`
	OperatingIp     string    `form:"operating_ip" json:"operating_ip" gorm:"comment:'IP'"`
	OperatingArea   string    `form:"operating_area" json:"operating_area" gorm:"comment:'归属地区'"`
	OperatingStatus string    `form:"operating_status" json:"operating_status" gorm:"comment:'成功/失败'"`
	MerchantTime    time.Time `form:"merchant_time" json:"merchant_time" gorm:"comment:'操作时间'"`
}

func init() {
	//自动迁移表,智慧添加缺少的字段,不会删除/更改缺少的字段
	_ = db.SuDB.AutoMigrate(&MerchantLog{})
}

func AddOperatingLog(types int, mid uint, name string, text string, ip string, area string, status string) (msg string, err error) {
	sudb := db.SuDB
	var m MerchantLog
	m.MerchantId = mid
	m.MerchantName = name
	m.OperatingText = text
	m.OperatingIp = ip
	m.OperatingArea = area
	m.OperatingStatus = status
	m.MerchantTime = time.Now()
	create := sudb.Create(&m)
	if create.Error != nil {
		return "", create.Error
	}
	return "成功", nil
}

func (m *MerchantLog) Create() (msg string, err error) {
	sudb := db.SuDB
	create := sudb.Create(&m)
	if create.Error != nil {
		return "", create.Error
	}
	return "成功", nil
}

func (m *MerchantLog) Retrieve() (msg string, err error) {
	panic("implement me")
}

func (m *MerchantLog) Update() (msg string, err error) {
	panic("implement me")
}

func (m *MerchantLog) Delect() (msg string, err error) {
	panic("implement me")
}
