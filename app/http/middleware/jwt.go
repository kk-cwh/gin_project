package jwt

import (
	"gin_project/lib/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/ginils"
)

func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.Query("token")
		code = 200
		if token == "" {
			code = 400
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 401
				default:
					code = 404
				}
			}
		}
		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}

}
