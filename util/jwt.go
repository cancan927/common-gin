package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

//签名使用的key
var stSignKey = []byte(viper.GetString("jwt.signKey"))

type JwtCustomClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, name string) (string, error) {
	iJwtCustomClaims := JwtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)

	//签名
	return token.SignedString(stSignKey)
}

func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	iJwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("invalid token")
	}

	return iJwtCustomClaims, err
}
