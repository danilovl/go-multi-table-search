package model

import (
	"net/http"
)

type ApiResponse interface {
	GetApiError() *ApiErrorType
	GetStatusCode() int
	Error() string
}

type ApiResponseType struct {
	ApiError *ApiErrorType          `json:"error"`
	Result   map[string]interface{} `json:"result"`
}

func (apiResponse ApiResponseType) GetApiError() *ApiErrorType {
	return apiResponse.ApiError
}

func (apiResponse ApiResponseType) Error() string {
	return apiResponse.ApiError.Error()
}

func (apiResponse ApiResponseType) GetStatusCode() int {
	if apiResponse.ApiError == nil {
		return http.StatusOK
	}

	return apiResponse.ApiError.GetStatusCode()
}
