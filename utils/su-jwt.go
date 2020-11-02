package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ilovesusu/su-gin/config"
	"time"
)

type CustomClaims struct {
	UID uint `json:"uid"` // 用户id
	MID uint `json:"mid"` // 商家id
	AID uint `json:"aid"` // 总后台id
	TID uint `json:"tid"` // 代理id
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

//token盐
func (j *JWT) NewJWT() *JWT {
	j.SigningKey = []byte(config.SuApp.JwtSecret)
	return j
}

//创建token
func (j *JWT) CreateAccessToken(claims *CustomClaims) (accesstoken string, err error) {
	claims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix() //过期时间
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accesstoken, err = token.SignedString(j.SigningKey)
	if err != nil {
		fmt.Println("refresh_token create error")
	}
	return accesstoken, nil
}

//验证token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func ParseToken(tokenString string) (claims *CustomClaims) {
	newJWT := JWT{}
	sujwt := newJWT.NewJWT()
	claims, _ = sujwt.ParseToken(tokenString) //解析token
	return
}
