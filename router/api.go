package router

import (
	"gin_project/app/http/controllers"
	"github.com/gin-gonic/gin"
)

/*
New 初始化路由
*/
func New(router *gin.Engine) {
	// 注册全局的中间件
	router.Use(gin.Logger(), gin.Recovery())
	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/categories", controllers.GetAllCategory)
		v1.POST("/categories", controllers.AddCategory)
		v1.PUT("/categories", controllers.UpdateCategory)
	}

}
