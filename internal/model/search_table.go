package model

import (
	"reflect"
)

type SearchTableType struct {
	Identifier    string   `json:"identifier"`
	Search        string   `json:"search"`
	TableName     string   `json:"tableName"`
	SelectColumns []string `json:"selectColumns"`
	WhereColumns  []string `json:"whereColumns"`
	Limit         int      `json:"limit"`
	Offset        int      `json:"offset"`
}

func (searchTableType SearchTableType) GetRequiredFieldMissing() []string {
	fields := [3]string{"Identifier", "TableName", "SelectColumns"}
	var missingFields []string

	for _, field := range fields {
		reflectValues := reflect.ValueOf(searchTableType)
		fieldValue := reflect.Indirect(reflectValues).FieldByName(field)

		if fieldValue.String() == "" || fieldValue.Len() == 0 {
			missingFields = append(missingFields, field)
		}

	}

	return missingFields
}
