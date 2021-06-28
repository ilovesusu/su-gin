package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/ilovesusu/su-gin/configs"
	"sync"
	"time"
)

var (
	once   sync.Once
	newJWT *JWT
	sujwt  = NewJWT()
)

type CustomClaims struct {
	//满足添加自定义参数需求
	jwt.StandardClaims
	UID int
}

type SuToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CustomClaimsOption func(*CustomClaims)

func WithUID(uid int) CustomClaimsOption {
	return func(c *CustomClaims) {
		c.UID = uid
	}
}

type JWT struct {
	SigningKey []byte
}

//Token 盐
func NewJWT() *JWT {
	once.Do(func() {
		newJWT = &JWT{}
		newJWT.SigningKey = []byte(configs.SuJwt.JwtSecret)
	})
	return newJWT
}

//创建 Token
func CreateAccessToken(options ...CustomClaimsOption) (accesstoken string, err error) {
	customClaims := &CustomClaims{}
	for _, option := range options {
		option(customClaims)
	}
	customClaims.ExpiresAt = time.Now().Add(configs.SuJwt.EXPIRES_AT_MAX_TIME * time.Hour).Unix() //过期时间
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	accesstoken, err = token.SignedString(sujwt.SigningKey)
	if err != nil {
		fmt.Println("refresh_token create error")
	}
	return accesstoken, nil
}

//验证 Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return sujwt.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
