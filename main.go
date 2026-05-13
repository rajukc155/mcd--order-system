package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Initialize DB
	initDB()
	defer db.Close()

	// Routing settings
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", orderHandler)

	// CORS settings (allows access from frontend)
	handler := cors.Default().Handler(mux)

	fmt.Println("Server starting: http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
