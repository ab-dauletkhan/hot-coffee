package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/handler/handler_utils"
	"github.com/ab-dauletkhan/hot-coffee/internal/service"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func PostInventory(w http.ResponseWriter, r *http.Request) {
	req := models.InventoryItem{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler_utils.ErrorResponseJSON(w, 400, "invalid request payload")
		service.CreateLog(
			r,
			slog.LevelError,
			http.StatusBadRequest,
			"invalid request payload",
		)
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
