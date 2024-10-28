package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	req := models.Order{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorJSONResponse(w, r, 400, "invalid request payload")
		return
	}
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
}

func PutOrder(w http.ResponseWriter, r *http.Request) {
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
}

func CloseOrder(w http.ResponseWriter, r *http.Request) {
}
