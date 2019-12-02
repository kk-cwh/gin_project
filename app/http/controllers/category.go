package controllers

import (
	"fmt"
	"gin_project/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
func AddCategory(c *gin.Context)  {
    //category := make(map[string]interface{})
	category := models.Category{}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
fmt.Println(category)
	if err:= models.AddCategory(&category);err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": category,
		})
	}
}

func UpdateCategory(c *gin.Context)  {

	category := &models.Category{}

	if err := c.ShouldBindJSON(category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//id, err1 :=  strconv.Atoi(c.Param("id"))
	//category.ID = 1
	//category.Name ="gogog"

	if err:= models.UpdateCategory(1,category);err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": category,
		})
	}
}
