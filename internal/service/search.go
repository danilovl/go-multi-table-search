package service

import (
	"github.com/gin-gonic/gin"
	"main/internal/helper"
	"main/internal/model"
	"main/internal/storage"
	"sync"
)

var SearchParam model.SearchParamType

func Search(context *gin.Context) map[string]interface{} {
	SearchParam = helper.InitSearchSearchParam(context)
	tableNames := SearchParam.GetTableNames()

	result := make(map[string]interface{})
	countTables := len(tableNames)

	var waitGroup sync.WaitGroup
	searchInTableChannel := make(chan model.SearchResultType, countTables)

	for _, searchTable := range SearchParam.SearchTables {
		waitGroup.Add(1)
		go searchInTable(&waitGroup, searchInTableChannel, searchTable)
	}

	waitGroup.Wait()
	close(searchInTableChannel)

	for searchTableResult := range searchInTableChannel {
		result[searchTableResult.Identifier] = searchTableResult.Result
	}

	return result
}

func searchInTable(
	waitGroup *sync.WaitGroup,
	channel chan model.SearchResultType,
	searchTable model.SearchTableType,
) {
	query := storage.GetQuery(searchTable)
	rows, err := storage.GetDB().Query(query)

	if err != nil {
		panic(err)
	}

	rowsCount := 0
	var tableResult []interface{}

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valueFields := make([]interface{}, count)

	for rows.Next() {
		rowsCount++

		for i := range columns {
			valueFields[i] = &values[i]
		}

		rows.Scan(valueFields...)
		rowResult := make(map[string]interface{})

		for i, column := range columns {
			val := values[i]

			bytes, ok := val.([]byte)
			var value interface{}
			if ok {
				value = string(bytes)
			} else {
				value = val
			}

			rowResult[column] = value
		}

		tableResult = append(tableResult, rowResult)
	}

	resultType := model.IdentifierResultType{Count: rowsCount, Result: tableResult}

	channel <- model.SearchResultType{Identifier: searchTable.Identifier, Result: resultType}
	waitGroup.Done()
}
