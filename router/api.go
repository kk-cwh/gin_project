package router

import (
	"gin_project/app/http/controllers"
	"gin_project/app/http/middleware"
	"github.com/gin-gonic/gin"
)

/*
InitRoutes 初始化路由
*/
func InitRoutes(router *gin.Engine) {
	// 注册全局的中间件
	router.Use(gin.Logger(), gin.Recovery(),middleware.TranslationMiddleware())

	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
	// 简单的路由组: v1
	v1 := router.Group("/v1")
	v1.Use(middleware.JwtAuth())
	{
		v1.POST("/articles", controllers.SaveArticle)
		v1.GET("/articles", controllers.GetArticles)
		v1.GET("/articles/:id", controllers.GetOneArticle)

		v1.GET("/categories", controllers.GetAllCategory)
		v1.POST("/categories", controllers.AddCategory)
		v1.PUT("/categories/:id", controllers.UpdateCategory)
	}

}
