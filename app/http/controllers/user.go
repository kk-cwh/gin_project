package controllers

import (
	"errors"
	"gin_project/app/models"
	"gin_project/lib/this"
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
	if err := this.BindingValidParams(c, loginInput); err != nil {
		this.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	dbPassword, err := util.ScryptStr(loginInput.Password)
	if err != nil {
		this.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	user, err := models.FindByUserName(loginInput.Username)
	if err != nil {
		this.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if user != nil && user.Password == dbPassword && user.State == 1 {
		claims := util.CreateCustomClaims(user.ID,user.Username)
		token, err := util.GenerateToken(&claims)
		this.Response(c, err, gin.H{"user_id":user.ID, "username": user.Username, "token": token})

	} else {
		this.ResponseError(c, http.StatusInternalServerError, errors.New("账号或密码有误"))
	}
	return

}

func Register(c *gin.Context) {
	inputUser := &models.User{}

	if err := this.BindingValidParams(c, inputUser); err != nil {
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
		this.ResponseError(c, http.StatusInternalServerError, errors.New("服务异常"))
		return
	}
	if user != nil {
		this.ResponseError(c, http.StatusBadRequest, errors.New("用户名已被使用"))
		return
	}

	err = models.SaveUser(inputUser)
	this.Response(c, err, inputUser)

}
