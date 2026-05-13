package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	if err := initDB("orders.db"); err != nil {
		log.Fatalf("データベースの初期化に失敗しました: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", createOrderHandler)

	handler := cors.Default().Handler(mux)

	log.Println("サーバー起動: http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("サーバー起動に失敗しました: %v", err)
	}
}
