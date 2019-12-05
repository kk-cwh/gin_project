package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type CustomClaims struct {
	UserID int `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
//生成自定义 Claims
func CreateCustomClaims(userId int ,username string)  (claims CustomClaims){
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims = CustomClaims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-project",
		},
	}
	return
}

// 生成token generate tokens used for auth
func GenerateToken(claims *CustomClaims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//解析token parsing token
func ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 更新token
func  RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
		return GenerateToken(claims)
	}
	return "", nil
}
