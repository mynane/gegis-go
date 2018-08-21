package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"github.com/astaxie/beego/logs"
	"fmt"
)

var (
	key []byte = []byte("-jwt-hzwy23@163.com")
)

func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 100),
		Issuer:    "hzwy23",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return ss
}

func CheckToken(token string) (bool, string) {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false, "认证失败"
	}

	return true, "认证成功"
}
