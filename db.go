package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	// Create/Open the SQLite database file
	db, err = sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}

	// Create table
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		menu TEXT,
		quantity INTEGER
	);`
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
}