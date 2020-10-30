package model

import (
	"encoding/json"
	"errors"
	"github.com/ilovesusu/su-gin/config"
	"io/ioutil"
	"os"
)

var (
	SuI18NList map[int]error
)

func init() {
	SuI18NList = make(map[int]error, 10)
	if config.SuI18n.I18N {
		//获取文件名
		filename := "i18n.Su." + config.SuApp.VERSION + "." + config.SuI18n.LANGUAGE
		if Exist(filename) {
			b, _ := ioutil.ReadFile(filename)
			_ = json.Unmarshal(b, &SuI18NList)
		} else {
			addError()
			addNotice()
			b, _ := json.MarshalIndent(&SuI18NList, "", "\t")
			_ = ioutil.WriteFile(filename, b, 0644)
		}
	}

}

//添加标准错误
func addError() {
	//系统错误
	SuI18NList[-1] = errors.New("SYSTEM ERROR")
	//统一成功的前缀
	SuI18NList[200] = errors.New("SUCCESS:")
	//统一失败的前缀
	SuI18NList[40000] = errors.New("ERROR:")
	SuI18NList[40001] = errors.New("PARAMETER ERROR:")
	SuI18NList[40002] = errors.New("LACK TOKEN:")
	SuI18NList[40003] = errors.New("Token ERROR:")
	SuI18NList[40004] = errors.New("Sign ERROR: Verify parameter failed")
	SuI18NList[40100] = errors.New("Gorm SQL Error:")
}

//添加标准通知
func addNotice() {
	SuI18NList[50000] = errors.New("Hello SuGin!")
	SuI18NList[50001] = errors.New("User Create Success!")
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
