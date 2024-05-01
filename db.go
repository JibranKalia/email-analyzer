package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open("sqlite3", "./email_data.db")
	if err != nil {
		return err
	}
	DB = db

	// Run migrations here...
	return nil
}
