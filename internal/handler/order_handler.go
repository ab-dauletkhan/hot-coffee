package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/handler/handler_utils"
	"github.com/ab-dauletkhan/hot-coffee/internal/service"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	req := models.Order{}
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
