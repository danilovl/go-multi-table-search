package main

import (
	"github.com/gin-gonic/gin"
	"main/internal/config"
	"main/internal/service"
)

func main() {
	config.Init()

	engine := gin.Default()
	engine.POST("/multi-table-search", func(context *gin.Context) {
		context.JSON(200, service.Search(context))
	})
	engine.Run()
}
