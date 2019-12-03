package controllers

import (
	"gin_project/app/models"
	"gin_project/lib/request"
	"gin_project/lib/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Username string `form:"username" json:"username" validate:"required" `
	Password string `form:"password" json:"password" validate:"required" `
}

func Login(c *gin.Context) {
	loginInput := &LoginInput{}
	if err := request.BindingValidParams(c, loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbPassword, err := util.ScryptStr(loginInput.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.FindByUserName(loginInput.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Password == dbPassword && user.State == 1 {
		claims := util.CreateCustomClaims(user.Username)
		token, err := util.GenerateToken(&claims)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"token":    token,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "error",
	})

}

func Register(c *gin.Context) {
	inputUser := &models.User{}

	if err := request.BindingValidParams(c, inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if dbPassword, err := util.ScryptStr(inputUser.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		inputUser.Password = dbPassword
		inputUser.State = 1
	}
	user, err := models.FindByUserName(inputUser.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}

	if err := models.SaveUser(inputUser); err == nil {
		c.JSON(http.StatusOK, gin.H{"error": "", "code": 0})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
