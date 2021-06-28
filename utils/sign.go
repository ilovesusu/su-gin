package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ilovesusu/su-gin/configs"
	"reflect"
	"sort"
	"strings"
)

type ParamM struct {
	Username  string `susu:"username"   form:"username"  json:"username" binding:"required"`
	Password  string `susu:"password"   form:"password"  json:"password" binding:"required"`
	SecretKey string `susu:"-" form:"secret_key" json:"tag_value reflect.Value" binding:"required"`
	TimeStamp int64  `susu:"time_stamp" form:"time_stamp" json:"time_stamp" binding:"required"`
}

type SignParam struct {
	Secret    string `form:"secret" susu:"-" json:"secret" binding:"required"`
	TimeStamp uint   `form:"time_stamp" susu:"time_stamp" json:"time_stamp" binding:"required"`
}

// 模拟微信计算签名
func Sign(a interface{}, tag string, key string) (sign string) {
	var (
		mReq      map[string]interface{}
		tag_value reflect.Value
	)
	//STEP0, 获取结构体 tag `xml or json` 的内容 作为 key 和 结构体参数的值为 value.
	valueOf := reflect.ValueOf(a)
	typeOf := reflect.TypeOf(a)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
		typeOf = typeOf.Elem()
	}
	mReq = make(map[string]interface{}, 0)
	for j := 0; j < valueOf.NumField(); j++ {
		tagname := typeOf.Field(j).Tag.Get(tag)
		//fmt.Println("tagname", tagname)
		if tagname != "-" && tagname != "" {
			tag_value = valueOf.Field(j)
			switch valueOf.Field(j).Kind() {
			case reflect.Int, reflect.Int64, reflect.Int32:
				mReq[tagname] = tag_value.Int()
			case reflect.Uint, reflect.Uint32, reflect.Uint64:
				mReq[tagname] = tag_value.Uint()
			case reflect.String:
				mReq[tagname] = tag_value.String()
			case reflect.Bool:
				mReq[tagname] = tag_value.Bool()
			}
		}
		if tagname == "" {
			sign := valueOf.Field(j).Interface().(SignParam)
			mReq["time_stamp"] = sign.TimeStamp
		}
	}
	//STEP1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)
	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	//if key != "" {
	//	signStrings = signStrings + "secret_key=" + key
	//}

	//STEP3,切除最后的&
	signStrings = signStrings[0 : len(signStrings)-1]

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	return upperSign
}

func CheckSign(a interface{}) bool {
	var tag_value_sign string
	sign := Sign(a, "susu", configs.SuApp.JwtSecret)
	valueOf := reflect.ValueOf(a)
	typeOf := reflect.TypeOf(a)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
		typeOf = typeOf.Elem()
	}

	for j := 0; j < valueOf.NumField(); j++ {
		tagname := typeOf.Field(j).Tag.Get("susu")
		//结构体tagname为空
		if tagname == "" {
			sign := valueOf.Field(j).Interface().(SignParam)
			tag_value_sign = sign.Secret
			break
		}
	}
	// 校验 sign
	//fmt.Println("tag_value_sign", tag_value_sign)
	if tag_value_sign == sign {
		return true
	}
	return false
}
