package config

import (
	"github.com/go-ini/ini"
	"time"
)

type app struct {
	APPNAME   string `ini:"APPNAME"`
	VERSION   string `ini:"VERSION"`
	PageSize  int    `ini:"PAGE_SIZE"`
	JwtSecret string `ini:"JWT_SECRET"`
	Run_Mode  string `ini:"RUN_MODE"`
	Assets    string `ini:"ASSETS"`
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

var (
	SuApp        app
	SuServer     server
	SuDatabase   database
	SuRedis      redis
	SuJwt        jwt
	SuI18n       i18n
	SuAlarmError alarmError
)

func init() {
	cfg, err := ini.Load("config/app.ini")
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
}
