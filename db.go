package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Order struct {
	ID       int64  `json:"id"`
	Menu     string `json:"menu"`
	Quantity int    `json:"quantity"`
}

var db *sql.DB

func initDB(path string) error {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        menu TEXT NOT NULL,
        quantity INTEGER NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`

	if _, err := db.Exec(createTableSQL); err != nil {
		return err
	}

	log.Println("SQLite データベースを初期化しました")
	return nil
}

func insertOrder(menu string, quantity int) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO orders(menu, quantity) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(menu, quantity)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
