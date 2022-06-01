package controller

import (
	"github.com/gin-gonic/gin"
	"main/internal/service"
)

type MultiTableSearchController struct {
}

func (controller MultiTableSearchController) Handler() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(200, service.Search(context))
	}
}

func NewProductController() MultiTableSearchController {
	return MultiTableSearchController{}
}
