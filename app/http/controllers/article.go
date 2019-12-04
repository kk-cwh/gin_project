package controllers

import (
	"gin_project/app/models"
	"gin_project/lib/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticles(c *gin.Context) {

	articles, err := models.GetArticles(0, 10)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{

			"Data": articles,
		})
	}

}
func GetOneArticle(c *gin.Context)  {
	id,_:= strconv.Atoi(c.Param("id"))

	article,err := models.GetOneArticle(id)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Data": article,
		})
	}

}

func SaveArticle(c *gin.Context) {
     article := &models.Article{}
	if err := request.BindingValidParams(c, article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.SaveArticle(article)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Data": article,
		})
	}

}
