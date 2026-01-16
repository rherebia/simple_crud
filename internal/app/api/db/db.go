package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	var err error
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	CreateTables(db)

	return db
}

func CreateTables(db *sql.DB) {
	createAlbumsTable := `
	CREATE TABLE IF NOT EXISTS albums (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		artist TEXT NOT NULL,
		price REAL NOT NULL
	)
	`

	_, err := db.Exec(createAlbumsTable)

	if err != nil {
		panic("Could not create albums table")
	}
}
