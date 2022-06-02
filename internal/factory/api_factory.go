package factory

import (
	"main/internal/constant"
	"main/internal/model"
)

func CreateApiErrorByType(errorType string) *model.ApiErrorType {
	var message []string
	message = append(message, constant.GetErrorTypeMessage(errorType))

	return &model.ApiErrorType{ErrorType: errorType, ErrorMessage: message}
}
