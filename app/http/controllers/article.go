package controllers

import (
	"gin_project/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {

	articles, err := models.GetArticles(0, 10)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{

			"Data": articles,
		})
	}

}
