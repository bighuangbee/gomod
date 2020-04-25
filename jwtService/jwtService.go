package jwtService

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/bighuangbee/gomod/config"
	redis2 "github.com/go-redis/redis"
	"github.com/bighuangbee/gomod/redis"
	"time"
)

/**
* @Author: bigHuangBee
* @Date: 2020/4/24 22:36
 */

const LOGIN_TOKEN_KEY = "sysUser:token_%s";
const INVALID_TOKEN_KEY = "sysUserTokenInvalid:%s"

type UserClaims struct {
	UserId uint `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
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

func GetExistsToken(userName string)(string, error){
	return redis.Redis.Get(CreateTokenKey(userName)).Result()
}

func SetToken(userName string, token string) error{
	return redis.Redis.Set(CreateTokenKey(userName), token, time.Minute * time.Duration(config.ConfigData.LoginExpire)).Err()
}

func DelToken(userName string){
	redis.Redis.Del(CreateTokenKey(userName))
}

// token索引键
func CreateTokenKey(userName string) (string){
	return fmt.Sprintf(LOGIN_TOKEN_KEY, userName)
}

// 加入到失效token列表
func JoinInvalidToken(token string){
	redis.Redis.Set(fmt.Sprintf(INVALID_TOKEN_KEY, token), token, time.Minute * time.Duration(config.ConfigData.LoginExpire))
}

// token是否已失效
func IsInvalidToken(token string) bool{
	_, err := redis.Redis.Get(fmt.Sprintf(INVALID_TOKEN_KEY, token)).Result()
	if err == redis2.Nil{
		return false
	}
	return true
}