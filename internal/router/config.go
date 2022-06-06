package router

import (
	"github.com/gin-gonic/gin"
	"main/internal/controller"
	"main/internal/middleware"
)

func Init(engine *gin.Engine) {
	engine.Use(middleware.SecurityHandler())
	engine.Use(middleware.ErrorHandler())
	engine.GET("/ping", controller.NewPingController().Handler())
	engine.POST("/multi-table-search", controller.NewMultiTableSearchController().Handler())
}
