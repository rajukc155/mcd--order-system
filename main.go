package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Initialize database
	InitDB()

	// API routes
	http.HandleFunc("/orders", OrdersHandler)

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(http.DefaultServeMux)

	log.Println("Server started: http://localhost:8080")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
