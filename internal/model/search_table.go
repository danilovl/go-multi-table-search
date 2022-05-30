package model

type SearchTableType struct {
	Identifier    string   `json:"identifier"`
	Search        string   `json:"search"`
	TableName     string   `json:"tableName"`
	SelectColumns []string `json:"selectColumns"`
	WhereColumns  []string `json:"whereColumns"`
	Limit         int      `json:"limit"`
	Offset        int      `json:"offset"`
}
