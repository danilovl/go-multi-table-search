package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"main/internal/constant"
	"main/internal/factory"
	"main/internal/model"
	"strings"
)

func InitSearchParam(context *gin.Context) (model.SearchParamType, model.ApiResponse) {
	SearchParam := model.SearchParamType{}

	searchTables, apiResponse := getSearchTables(context)
	SearchParam.SearchTables = searchTables

	return SearchParam, apiResponse
}

func getSearchTables(context *gin.Context) ([]model.SearchTableType, model.ApiResponse) {
	var unmarshalSearchTableTypes []model.SearchTableType
	var searchTableTypes []model.SearchTableType

	body, errReadBody := ioutil.ReadAll(context.Request.Body)
	if errReadBody != nil {
		apiResponse := model.ApiResponseType{
			ApiError: factory.CreateApiErrorByType(constant.ReadError),
		}

		return nil, apiResponse
	}

	errUnmarshal := json.Unmarshal(body, &unmarshalSearchTableTypes)
	if errUnmarshal != nil {
		apiResponse := model.ApiResponseType{
			ApiError: factory.CreateApiErrorByType(constant.UnmarshalError),
		}

		return nil, apiResponse
	}

	var missingFieldsErrorMassages []string

	for _, searchTableType := range unmarshalSearchTableTypes {
		missingFields := searchTableType.GetRequiredFieldMissing()

		if len(missingFields) > 0 {
			missingFieldsErrorMassages = append(
				missingFieldsErrorMassages,
				"Please control require fields: "+strings.Join(missingFields[:], ","),
			)

			continue
		}

		if searchTableType.Identifier == "" || searchTableType.TableName == "" || len(searchTableType.SelectColumns) == 0 {
			continue
		}

		if searchTableType.SelectColumns[0] != "*" && (searchTableType.Search == "" || len(searchTableType.WhereColumns) == 0) {
			continue
		}

		limit := searchTableType.Limit
		offset := searchTableType.Offset

		switch {
		case limit > 500:
			limit = 500
		case limit <= 0:
			limit = 10
		}

		switch {
		case offset < 0:
			offset = 0
		}

		searchTableType.Limit = limit
		searchTableType.Offset = offset

		searchTableTypes = append(searchTableTypes, searchTableType)
	}

	if len(missingFieldsErrorMassages) > 0 {
		apiResponse := model.ApiResponseType{
			ApiError: &model.ApiErrorType{ErrorType: constant.ParamError, ErrorMessage: missingFieldsErrorMassages},
		}

		return nil, apiResponse
	}

	return searchTableTypes, nil
}
