package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	IsAdmin  int    `json:"is_admin"`
	jwt.StandardClaims
}

var MyKey = []byte("hello")

func (c *UserClaims) GenerateToken() (string, error) {
	// 根据claims生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 根据自定义key生成tokenString
	return token.SignedString(MyKey)
}

func (c *UserClaims) ParseToken(str string) error {
	token, err := jwt.ParseWithClaims(str, c, func(token *jwt.Token) (interface{}, error) {
		return MyKey, nil
	})
	if token.Valid != true {
		return errors.New("token不合法")
	}
	return err
}
