package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

func PostInventory(w http.ResponseWriter, r *http.Request) {
	req := models.InventoryItem{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponseJSON(w, 400, "Invalid request payload.")
		return
	}
}

func GetAllInventory(w http.ResponseWriter, r *http.Request) {
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
}

func PutInventory(w http.ResponseWriter, r *http.Request) {
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
}
