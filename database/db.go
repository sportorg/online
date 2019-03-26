package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func Connect(dataSourceName string) *sql.DB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	return db
}