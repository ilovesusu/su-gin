package utils

import (
	"fmt"
	"github.com/ilovesusu/su-gin/configs"
	"runtime/debug"
	"strings"
	"time"
)

type errorString struct {
	s string
}

// ErrorMail 发邮件
func ErrorMail(text string) bool {
	DebugStack := ""
	for _, v := range strings.Split(string(debug.Stack()), "\n") {
		DebugStack += v + "<br>"
	}

	subject := fmt.Sprintf("【系统告警】%s 项目出错了！", configs.SuApp.APPNAME)

	body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", text))
	body = strings.ReplaceAll(body, "{RequestTime}", time.Now().Format("2006/01/02 15:04:05"))
	body = strings.ReplaceAll(body, "{RequestURL}", "-URL-")
	body = strings.ReplaceAll(body, "{RequestUA}", "-UA-")
	body = strings.ReplaceAll(body, "{RequestIP}", "-IP-")
	body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

	// 执行发邮件
	options := &Options{
		MailHost: configs.SuAlarmError.EmailHost,
		MailPort: configs.SuAlarmError.EmailPort,
		MailUser: configs.SuAlarmError.EmailUser,
		MailPass: configs.SuAlarmError.EmailPass,
		MailTo:   configs.SuAlarmError.ErrorNotifyUser,
		Subject:  subject,
		Body:     body,
	}
	err := SendEmail(options)
	if err != nil {
		return false
	}
	return true
}

// 发短信
func ErrorSms(text string) bool {
	return true
}

// 发微信
func ErrorWeChat(text string) bool {
	return true
}
