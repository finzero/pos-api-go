package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDB() (*sql.DB, error) {
	// initialize connection to DB
	var err error
	DB, err = sql.Open("sqlite3", "./posapi.db")
	return DB, err
}
