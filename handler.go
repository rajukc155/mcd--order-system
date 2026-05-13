package main

import (
	"encoding/json"
	"net/http"
	"database/sql"
)

type Order struct {
	ID       int    `json:"order_id"`
	Menu     string `json:"menu"`
	Quantity int    `json:"quantity"`
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var ord Order
	if err := json.NewDecoder(r.Body).Decode(&ord); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save to DB
	res, err := db.Exec("INSERT INTO orders (menu, quantity) VALUES (?, ?)", ord.Menu, ord.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "success",
		"order_id": id,
		"message":  "Order received",
	})
}
)


