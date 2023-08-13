package server

import (
	"database/sql"
	"nicetry/helper"
)

func NewDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/nicetry")
	helper.PanicIfError(err)

	return db
}
