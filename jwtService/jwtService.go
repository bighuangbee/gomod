package jwtService

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/bighuangbee/gomod/config"
)

/**
* @Author: bigHuangBee
* @Date: 2020/4/24 22:36
 */

type UserClaims struct {
	UserId uint `json:"user_id"`
	Roles []string `json:"roles"`
	Uuid string `json:"deivce_uuid"`
	jwt.StandardClaims
}

func CreateToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.ConfigData.JwtEncrtpy))
}

func ParseToken(tokenStr string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.ConfigData.JwtEncrtpy), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}


	return nil, err
}