package this

import "github.com/gin-gonic/gin"

type User struct {
	Id int
	Username string
}

func Auth(c *gin.Context) *User{
	user :=&User{}
	user.Id = c.GetInt("user_id")
	user.Username = c.GetString("username")
	return user
}
