package controllers

import (
	"errors"
	"gin_project/app/models"
	"gin_project/lib/this"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllCategory all
func GetAllCategory(c *gin.Context) {
	articles, err := models.GetAllCategory()
	this.Response(c,err,articles)
	return

}

// AddCategory add
func AddCategory(c *gin.Context) {
	loginInput := models.Category{}
	err := this.BindingValidParams(c, &loginInput)
	if err != nil {
		this.ResponseError(c,http.StatusBadRequest,err)
		return
	}
	err = models.AddCategory(&loginInput)
	this.Response(c,err,loginInput)
	return
}

func UpdateCategory(c *gin.Context) {
	category := models.Category{}
	if err := this.BindingValidParams(c,&category); err != nil {
		this.ResponseError(c,http.StatusBadRequest,err)
		return
	}

	id,err:= strconv.Atoi(c.Param("id"))
	if err!=nil {
		this.ResponseError(c,http.StatusBadRequest,errors.New("id参数有误"))
		return
	}

	 err = models.UpdateCategory(id, &category)
	 this.Response(c,err,category)
	 return

}

