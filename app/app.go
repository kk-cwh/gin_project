package app

import (
	"gin_project/router"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	engine := gin.New()
	// 配置路由
	router.New(engine)

	return engine
}
