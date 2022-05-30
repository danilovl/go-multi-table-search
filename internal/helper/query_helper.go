package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"main/internal/model"
)

func InitSearchSearchParam(context *gin.Context) model.SearchParamType {
	SearchParam := model.SearchParamType{}
	SearchParam.SearchTables = getSearchTables(context)

	return SearchParam
}

func getSearchTables(context *gin.Context) []model.SearchTableType {
	var unmarshalSearchTableTypes []model.SearchTableType
	var searchTableTypes []model.SearchTableType

	body, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &unmarshalSearchTableTypes)

	for _, searchTableType := range unmarshalSearchTableTypes {
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

	return searchTableTypes
}
