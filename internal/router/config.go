package router

import (
	"github.com/gin-gonic/gin"
	"main/internal/controller"
	"main/internal/middleware"
)

func Init(engine *gin.Engine) {
	engine.Use(middleware.ErrorHandler())
	engine.POST("/multi-table-search", controller.NewProductController().Handler())
}
