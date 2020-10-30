package utils

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	var susu ParamM
	susu.Username = "susu"
	susu.Password = "9527"
	susu.TimeStamp = time.Now().Unix()
	sign := Sign(&susu, "susu", "susuwoaini")
	fmt.Println(sign)

	susu.SecretKey = sign

	checkSign := CheckSign(susu)
	fmt.Println(checkSign)

	fmt.Println(reflect.TypeOf(ParamM{}).Name())
}
