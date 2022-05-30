package model

type IdentifierResultType struct {
	Count  int           `json:"count"`
	Result []interface{} `json:"result"`
}

type SearchResultType struct {
	Identifier string
	Result     IdentifierResultType
}
