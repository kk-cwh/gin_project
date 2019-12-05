package middleware

import (
	"gin_project/lib/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.Request.Header.Get("Authorization")
		code = 200
		if token == "" {
			code = 400
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 401
				default:
					code = 404
				}
			}else {
				c.Set("user_id", claims.UserID)
				c.Set("username", claims.Username)
			}
		}
		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "token认证无效",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
