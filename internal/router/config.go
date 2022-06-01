package router

import (
	"github.com/gin-gonic/gin"
	"main/internal/controller"
)

func Init(engine *gin.Engine) {
	engine.POST("/multi-table-search", controller.NewProductController().Handler())
}
