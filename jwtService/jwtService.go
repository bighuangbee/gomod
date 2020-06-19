package jwtService

import (
	"fmt"
	"github.com/bighuangbee/gomod/config"
	"github.com/bighuangbee/gomod/http/respone"
	"github.com/bighuangbee/gomod/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-redis/redis"
	"time"
)

/**
* @Author: bigHuangBee
* @Date: 2020/4/24 22:36
 */

const JWT_KEY_SYS = "sysUser"
const JWT_KEY_DUTY = "dutyUser"
const JWT_KEY_DRONE = "droneUser"

const USER_TYPE_MOINTOR = 1
const USER_TYPE_DUTY = 2
const USER_TYPE_DRONE = 3

type UserClaims struct {
	UserId uint `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Roles []string `json:"roles"`
	Uuid string `json:"deivce_uuid"`
	Type int `json:"type"`
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

func (user *UserJwt)SetToken(userName string, token string) error{
	return redis.Redis.Set(user.CreateTokenKey(userName), token, time.Minute * time.Duration(config.ConfigData.LoginExpire)).Err()
}

func NewUserJwt(userType string) *UserJwt{
	return &UserJwt{
		Type:            userType,
		Encrtpy:         userType + config.ConfigData.JwtEncrtpy,
		TokenKey:        userType + ":token_%s",
		InValidTokenKey: userType + "TokenInvalid:%s",
	}
}

func (user *UserJwt)GetExistsToken(userName string)(string, error){
	return redis.Redis.Get(user.CreateTokenKey(userName)).Result()
}

func (user *UserJwt)DelToken(userName string) error{
	return redis.Redis.Del(user.CreateTokenKey(userName)).Err()
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
	token, err := redis.Redis.Get(fmt.Sprintf(user.InValidTokenKey, token)).Result()
	if err == redis2.Nil{
		return false
	}
	return true
}



func Authorization(userType string) gin.HandlerFunc{

	return func(c *gin.Context) {

		tokenStr := c.GetHeader("Authorization")
		if tokenStr == ""{
			respone.UnAuthorized(c, "登录令牌为空")
			return
		}

		jwt := NewUserJwt(userType)
		claims, err := jwt.ParseToken(tokenStr)
		if err == nil {
			existsToken, err := jwt.GetExistsToken(claims.UserName)
			if err != nil || existsToken == ""{
				respone.UnAuthorized(c, "您的帐户登录失效")
				return
			}

			if jwt.IsInvalidToken(tokenStr) {
				respone.UnAuthorized(c, "您的帐户异地登或令牌失效")
				return
			}

			c.Set("user_id", int64(claims.UserId))
			c.Set("roles", claims.Roles)
			c.Set("device_uuid", claims.Uuid)

			c.Set("uuid", claims.Uuid)
			c.Set("claims", claims)

			c.Next()
			return
		}

		respone.UnAuthorized(c,err.Error())
	}
}