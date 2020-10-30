package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//md5加密接口
func SuMd5(salt ...string) (data string, err error) {
	if len(salt) > 0 {
		mima := salt[0]
		if len(salt) > 1 && salt[1] != "" {
			bytes := []byte(mima + salt[1])
			sum := md5.Sum(bytes)
			return fmt.Sprintf("%x", sum), nil
		} else {
			bytes := []byte(mima)
			sum := md5.Sum(bytes)
			return fmt.Sprintf("%x", sum), nil
		}
	} else {
		return data, errors.New("数据为空")
	}
}

// 随机生成指定位数的大写字母和数字的组合
func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
