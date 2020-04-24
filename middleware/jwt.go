package middleware

import (
	"github.com/bighuangbee/air-land-command/pkg/message"
	"github.com/bighuangbee/air-land-command/pkg/respone"
	"github.com/bighuangbee/gomod/jwtService"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc{

	return func(c *gin.Context) {

		tokenStr := c.GetHeader("Authorization")
		if tokenStr == ""{
			respone.UnAuthorized(c, message.USER_AUTHORIZATION_EMPTY)
			return
		}

		claims, err := jwtService.ParseToken(tokenStr)

		if err == nil {
			c.Set("user_id", int64(claims.UserId))
			c.Set("roles", claims.Roles)
			c.Set("uuid", claims.Uuid)

			c.Next()
			return
		}

		respone.UnAuthorized(c,err.Error())
	}
}