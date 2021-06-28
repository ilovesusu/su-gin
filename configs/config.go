package configs

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"time"
)

type app struct {
	APPNAME   string `ini:"APPNAME"`
	VERSION   string `ini:"VERSION"`
	PageSize  int    `ini:"PAGE_SIZE"`
	PWDSecret string `ini:"PWD_SECRET"`
	Run_Mode  string `ini:"RUN_MODE"`
	Assets    string `ini:"ASSETS"`
	Log       string `ini:"LOG"`
	SaveLog   bool   `ini:"SAVE_LOG"`
}

type server struct {
	Domain       string        `ini:"DOMAIN_NAME"`
	HttpIP       string        `ini:"HTTP_IP"`
	HttpPort     string        `ini:"HTTP_PORT"`
	ReadTimeout  time.Duration `ini:"READ_TIMEOUT"`
	WriteTimeout time.Duration `ini:"WRITE_TIMEOUT"`
}

type database struct {
	Types              string        `ini:"TYPES"`
	User               string        `ini:"USER"`
	PassWord           string        `ini:"PASSWORD"`
	Host               string        `ini:"HOST"`
	Port               string        `ini:"PORT"`
	DatabaseName       string        `ini:"DATABASE_NAME"`
	TablePrefix        string        `ini:"TABLE_PREFIX"`
	Log_Mode           bool          `ini:"LOG_MODE"`
	MAX_IDLE_CONNS     int           `ini:"MAX_IDLE_CONNS"`
	MAX_OPEN_CONNS     int           `ini:"MAX_OPEN_CONNS"`
	CONN_MAX_LIFE_TIME time.Duration `ini:"CONN_MAX_LIFE_TIME"`
	SINGULAR_TABLE     bool          `ini:"SINGULAR_TABLE"`
}

type redis struct {
	Types          string `ini:"TYPES"`
	PassWord       string `ini:"PASSWORD"`
	Host           string `ini:"HOST"`
	Port           string `ini:"PORT"`
	DatabasesName  int    `ini:"DATABASES_NAME"`
	POOL_SIZE      int    `ini:"POOL_SIZE"`
	MIN_IDLE_CONNS int    `ini:"MIN_IDLE_CONNS"`
}

type jwt struct {
	EXPIRES_AT_MAX_TIME time.Duration `ini:"EXPIRES_AT_MAX_TIME"`
	JwtSecret           string        `ini:"JWT_SECRET"`
	JwtResetSecret      string        `ini:"JWT_RESET_SECRET"`
}

type i18n struct {
	I18N     bool   `ini:"I18N"`
	LANGUAGE string `ini:"LANGUAGE"`
}

type alarmError struct {
	EmailHost       string `ini:"EMAIL_HOST"`
	EmailPort       int    `ini:"EMAIL_PORT"`
	EmailUser       string `ini:"EMAIL_USER"`
	EmailPass       string `ini:"EMAIL_PASS"`
	ErrorNotifyUser string `ini:"ERROR_NOTIFY_USER"`
	UserSMS         string `ini:"USER_SMS"`
	UserWX          string `ini:"USER_WX"`
	ErrorNotifyOpen int    `ini:"ERROR_NOTIFY_OPEN"`
}

type aliSms struct {
	RegionId     string `ini:"RegionId"`
	KeyId        string `ini:"KeyId"`
	KeySecret    string `ini:"KeySecret"`
	SignName     string `ini:"SignName"`
	LoginCode    string `ini:"LoginCode"`
	RegisterCode string `ini:"RegisterCode"`
	TaskCode     string `ini:"TaskCode"`
	ResCode      string `ini:"ResCode"`
}

type juheSms struct {
	TplId string `ini:"TplId"`
	Key   string `ini:"Key"`
}

var (
	SuApp        app
	SuServer     server
	SuDatabase   database
	SuRedis      redis
	SuJwt        jwt
	SuI18n       i18n
	SuAlarmError alarmError
	SuAliSms     aliSms
	SuJuheSms    juheSms
)

func init() {
	cfg, err := ini.Load("configs/app.ini")
	if err != nil {
		panic(err)
	}

	err = cfg.Section("app").MapTo(&SuApp)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("server").MapTo(&SuServer)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("database").MapTo(&SuDatabase)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("redis").MapTo(&SuRedis)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("jwt").MapTo(&SuJwt)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("i18n").MapTo(&SuI18n)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("alarm").MapTo(&SuAlarmError)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("ali_sms").MapTo(&SuAliSms)
	if err != nil {
		panic(err)
	}

	err = cfg.Section("juhe_sms").MapTo(&SuJuheSms)
	if err != nil {
		panic(err)
	}

	InitFileCreate(&SuApp)
}

//初始化项目文件夹
func InitFileCreate(suApp *app) {
	// 资源目录
	err := FileMkdir(suApp.Assets)
	if !err {
		os.Exit(0)
	}
	// 导出资源目录
	err = FileMkdir(suApp.Assets + "/file")
	if !err {
		os.Exit(0)
	}
	// 日志目录
	err = FileMkdir(suApp.Log)
	if !err {
		os.Exit(0)
	}
}

//创建新文件夹
func FileMkdir(filename string) bool {
	err := os.MkdirAll(filename, 0755)
	if os.IsExist(err) {
		s := fmt.Sprintf("文件夹%s创建失败", filename)
		fmt.Println(s)
		return false
	} else {
		return true
	}
}
