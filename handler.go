package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type createOrderRequest struct {
	Menu     string `json:"menu"`
	Quantity int    `json:"quantity"`
}

type createOrderResponse struct {
	Status  string `json:"status"`
	OrderID int64  `json:"order_id,omitempty"`
	Message string `json:"message"`
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSON(w, http.StatusMethodNotAllowed, createOrderResponse{
			Status:  "error",
			Message: "POSTのみ対応しています",
		})
		return
	}

	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		sendJSON(w, http.StatusBadRequest, createOrderResponse{
			Status:  "error",
			Message: "Content-Typeはapplication/jsonで指定してください",
		})
		return
	}

	var req createOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSON(w, http.StatusBadRequest, createOrderResponse{
			Status:  "error",
			Message: "JSONの解析に失敗しました",
		})
		return
	}

	if req.Menu == "" || req.Quantity <= 0 {
		sendJSON(w, http.StatusBadRequest, createOrderResponse{
			Status:  "error",
			Message: "menuとquantityを正しく指定してください",
		})
		return
	}

	id, err := insertOrder(req.Menu, req.Quantity)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, createOrderResponse{
			Status:  "error",
			Message: "注文の保存に失敗しました",
		})
		return
	}

	sendJSON(w, http.StatusOK, createOrderResponse{
		Status:  "success",
		OrderID: id,
		Message: "注文を受け付けました",
	})
}

func sendJSON(w http.ResponseWriter, status int, payload createOrderResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
