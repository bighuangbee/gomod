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

type UserClaims struct {
	UserId uint `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Roles []string `json:"roles"`
	Uuid string `json:"deivce_uuid"`
	jwt.StandardClaims
}

type UserJwt struct{
	Type 		string	//用户类型
	Encrtpy 	string	//密钥
	TokenKey 	string	//token索引
	InValidTokenKey string	//失效token索引
}

func (user *UserJwt)CreateToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(user.Encrtpy))
}

func (user *UserJwt)ParseToken(tokenStr string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(user.Encrtpy), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}


	return nil, err
}

func (user *UserJwt)GetExistsToken(userName string)(string, error){
	return redis.Redis.Get(user.CreateTokenKey(userName)).Result()
}

func (user *UserJwt)SetToken(userName string, token string) error{
	return redis.Redis.Set(user.CreateTokenKey(userName), token, time.Minute * time.Duration(config.ConfigData.LoginExpire)).Err()
}

func (user *UserJwt)DelToken(userName string){
	redis.Redis.Del(user.CreateTokenKey(userName))
}

// token索引键
func (user *UserJwt)CreateTokenKey(userName string) (string){
	return fmt.Sprintf(user.TokenKey, userName)
}

// 加入到失效token列表
func (user *UserJwt)JoinInvalidToken(token string){
	redis.Redis.Set(fmt.Sprintf(user.InValidTokenKey, token), token, time.Minute * time.Duration(config.ConfigData.LoginExpire))
}

// token是否已失效
func (user *UserJwt)IsInvalidToken(token string) bool{
	_, err := redis.Redis.Get(fmt.Sprintf(user.InValidTokenKey, token)).Result()
	if err == redis2.Nil{
		return false
	}
	return true
}