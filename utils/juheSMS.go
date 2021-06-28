package utils

import (
	"fmt"
	"github.com/ilovesusu/su-gin/configs"
	"io/ioutil"
	"net/http"
	"net/url"
)

// todo 聚合短信
func JuheCurl(code string, mobile string) (string, error) {
	//请求地址
	juheURL := "http://v.juhe.cn/sms/send"

	//初始化参数
	param := url.Values{}
	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("mobile", mobile)                              //接受短信的用户手机号码
	param.Set("tpl_id", configs.SuJuheSms.TplId)              //您申请的短信模板ID，根据实际情况修改
	param.Set("tpl_value", "#code#="+code+"&#company#=聚合数据") //您设置的模板变量，根据实际情况
	param.Set("key", configs.SuJuheSms.Key)                   //应用APPKEY(应用详细页查询)

	Url, erk := url.Parse(juheURL)
	if erk != nil {
		fmt.Printf("解析url错误:\r\n%v", erk)
		return "", erk
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = param.Encode()
	request, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}

	defer request.Body.Close()
	content, ers := ioutil.ReadAll(request.Body)
	if ers != nil {
		fmt.Println("Fatal error ", ers.Error())
		return "", ers
	}
	return string(content), nil
}
