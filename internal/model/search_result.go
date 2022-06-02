package model

type IdentifierResultType struct {
	Count  int           `json:"count"`
	Result []interface{} `json:"result"`
	Error  interface{}   `json:"error"`
}

type SearchResultType struct {
	Identifier string
	Result     IdentifierResultType
}
