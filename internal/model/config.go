package model

import (
	"database/sql"
)

type ConfigType struct {
	Storage *sql.DB
}
