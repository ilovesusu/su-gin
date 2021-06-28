package utils

import (
	"fmt"
	"github.com/ilovesusu/su-gin/configs"
	"reflect"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	var susu ParamM
	susu.Username = "susu"
	susu.Password = "9527"
	susu.TimeStamp = time.Now().Unix()
	sign := Sign(&susu, "susu", configs.SuApp.JwtSecret)

	susu.SecretKey = sign

	checkSign := CheckSign(susu)
	fmt.Println(checkSign)

	fmt.Println(reflect.TypeOf(ParamM{}).Name())
}
