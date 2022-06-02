package model

import (
	"main/internal/constant"
	"net/http"
)

type ApiError interface {
	GetStatusCode() int
}

type ApiErrorType struct {
	ErrorType    string   `json:"type"`
	ErrorMessage []string `json:"message"`
}

func (apiError ApiErrorType) GetStatusCode() int {
	switch apiError.ErrorType {
	case constant.ReadError, constant.ParamError, constant.UnmarshalError:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func (apiError ApiErrorType) Error() string {
	return constant.GetErrorTypeMessage(apiError.ErrorType)
}
