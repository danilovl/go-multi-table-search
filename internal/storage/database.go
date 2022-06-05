package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Initialize(dsn string) {
	connection, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	db = connection
}

func GetDB() *sql.DB {
	return db
}
