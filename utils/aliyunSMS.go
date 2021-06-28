package utils

import (
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/ilovesusu/su-gin/configs"
)

// todo aliyun短信
func AliyunSms(code string, mobile string, status string) (msg string, err error) {
	print(status)
	client, err := dysmsapi.NewClientWithAccessKey(configs.SuAliSms.RegionId,
		configs.SuAliSms.KeyId,
		configs.SuAliSms.KeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = mobile
	request.SignName = configs.SuAliSms.SignName
	switch status {
	//登录
	case "login":
		request.TemplateCode = configs.SuAliSms.LoginCode
		request.TemplateParam = "{\"code\":\"" + code + "\"}"
	//注册
	case "reg":
		request.TemplateCode = configs.SuAliSms.RegisterCode
		request.TemplateParam = "{\"code\":\"" + code + "\"}"
	//接单
	case "task":
		request.TemplateCode = configs.SuAliSms.TaskCode
		//request.TemplateParam = "{\"status\":\"待操作\",\"time\":\"" + code + "\"}"
		request.TemplateParam = "{\"code\":\"" + code + "\"}"
	//重置密码
	case "res":
		request.TemplateCode = configs.SuAliSms.ResCode
		request.TemplateParam = "{\"code\":\"" + code + "\"}"
	}

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	if response.Code == "OK" {
		msg = "发送短信成功"
		return
	} else {
		msg = "发送短信失败"
		err = errors.New("发送失败")
		return
	}
}
