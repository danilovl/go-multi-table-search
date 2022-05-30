package model

type SearchParamType struct {
	SearchTables []SearchTableType
}

func (searchParamType SearchParamType) GetTableNames() []string {
	var result []string

	for _, table := range searchParamType.SearchTables {
		result = append(result, table.TableName)
	}

	return result
}
