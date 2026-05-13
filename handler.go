package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Order struct {
	ID       int    `json:"id"`
	Menu     string `json:"menu"`
	Quantity int    `json:"quantity"`
}

type Response struct {
	Status  string `json:"status"`
	OrderID int64  `json:"order_id,omitempty"`
	Message string `json:"message"`
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		CreateOrder(w, r)

	case http.MethodGet:
		GetOrders(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"INSERT INTO orders(menu, quantity) VALUES(?, ?)",
		order.Menu,
		order.Quantity,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()

	response := Response{
		Status:  "success",
		OrderID: id,
		Message: "Order received successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, menu, quantity FROM orders")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order

		err := rows.Scan(&order.ID, &order.Menu, &order.Quantity)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		orders = append(orders, order)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
