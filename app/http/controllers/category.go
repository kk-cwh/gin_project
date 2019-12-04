package controllers

import (
	"fmt"
	"gin_project/app/models"
	"gin_project/lib/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllCategory all
func GetAllCategory(c *gin.Context) {
	articles, err := models.GetAllCategory()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": articles,
		})
	}

}

// AddCategory add
func AddCategory(c *gin.Context) {
	loginInput := models.Category{}
	err := request.BindingValidParams(c, &loginInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//category := models.Category{}
	//if err := c.ShouldBindJSON(&category); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//fmt.Println(category)
		if err:= models.AddCategory(&loginInput);err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": loginInput,
			})
		}
}

func UpdateCategory(c *gin.Context) {

	category := models.Category{}
	if err := request.BindingValidParams(c,&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id,_:= strconv.Atoi(c.Param("id"))

	if err := models.UpdateCategory(id, &category); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": category,
		})
	}else {
		fmt.Println(err.Error())
	}

}

