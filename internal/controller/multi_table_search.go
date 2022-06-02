package controller

import (
	"github.com/gin-gonic/gin"
	"main/internal/service"
)

type MultiTableSearchController struct{}

func (controller MultiTableSearchController) Handler() func(context *gin.Context) {
	return func(context *gin.Context) {
		apiResponse := service.Search(context)

		if apiResponse.GetApiError() != nil {
			context.Error(apiResponse)

			return
		}

		context.JSON(apiResponse.GetStatusCode(), apiResponse)
	}
}

func NewProductController() MultiTableSearchController {
	return MultiTableSearchController{}
}
