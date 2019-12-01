package app

import (
	"gin_project/app/models"
	"gin_project/lib/setting"
	"gin_project/router"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	setting.Setup()
	models.Setup()
	engine := gin.New()
	// 配置路由
	router.New(engine)

	return engine
}
