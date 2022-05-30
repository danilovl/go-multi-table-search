package storage

import (
	"fmt"
	"main/internal/model"
)

func GetQuery(searchTable model.SearchTableType) string {
	query := "SELECT "
	lenSelectColumns := len(searchTable.SelectColumns)

	for index, selectColumn := range searchTable.SelectColumns {
		suffix := " "
		if lenSelectColumns > 1 && index < lenSelectColumns-1 {
			suffix = ", "
		}

		query += selectColumn + suffix
	}

	query += fmt.Sprintf("FROM %s", searchTable.TableName)

	for i, tableColumn := range searchTable.WhereColumns {
		orWhere := " "
		if i > 1 {
			orWhere = "OR"
		}

		query += fmt.Sprintf("%s WHERE %s LIKE '%%%s%%'", orWhere, tableColumn, searchTable.Search)
	}

	query += " ORDER BY id"
	query += fmt.Sprintf(" LIMIT %d", searchTable.Limit)
	query += fmt.Sprintf(" OFFSET %d;", searchTable.Offset)

	return query
}
