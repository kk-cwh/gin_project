package controllers

import (
	"gin_project/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	articles, err := models.GetArticles(0, 10)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{

			"Data": articles,
		})
	}

}
