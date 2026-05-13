package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error

	db, err = sql.Open("sqlite3", "./orders.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		menu TEXT,
		quantity INTEGER
	);
	`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialized")
}
