package middleware

import (
	"github.com/gin-gonic/gin"
	"main/internal/model"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		for _, apiResponse := range context.Errors {
			switch apiResponse.Err.(type) {
			case model.ApiResponseType:
				statusCode := apiResponse.Err.(model.ApiResponseType).GetStatusCode()
				response := apiResponse.Err.(model.ApiResponseType)

				context.JSON(statusCode, response)
			default:
				context.JSON(-1, apiResponse.Error())
			}
		}
	}
}
