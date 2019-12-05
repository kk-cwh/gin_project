package controllers

import (
	"gin_project/app/models"
	"gin_project/lib/this"
	"gin_project/lib/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type  InputArticle struct {
	Title        string   `json:"title" validate:"required" `
	ContentMd    string   `json:"content_md" validate:"required" `
	ContentHTML  string   `json:"content_html" validate:"required" `
	PageImageURL string   `json:"page_image_url" validate:"required" `
	CategoryID   uint     `json:"category_id" validate:"required" `
	State        int      `json:"state" validate:"required" `
	PageView     int      `json:"page_view"  `
}

func GetArticles(c *gin.Context) {

	articles, err := models.GetArticles(0, 10)
	this.Response(c,err,articles)
   return

}
func GetOneArticle(c *gin.Context)  {
	id,_:= strconv.Atoi(c.Param("id"))
	article,err := models.GetOneArticle(id)
	this.Response(c,err,article)
	return

}

func SaveArticle(c *gin.Context) {
     inputArticle := &InputArticle{}
	if err := this.BindingValidParams(c, inputArticle); err != nil {
		this.ResponseError(c,http.StatusBadRequest,err)
		return
	}
	article := &models.Article{}
	util.CopyStruct(inputArticle,article)
	article.UserID = this.Auth(c).Id
	err := models.SaveArticle(article)
	this.Response(c,err,article)
	return

}
