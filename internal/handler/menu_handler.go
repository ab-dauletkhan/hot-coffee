package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/handler/handler_utils"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func PostMenu(w http.ResponseWriter, r *http.Request) {
	req := models.MenuItem{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler_utils.ErrorJSONResponse(w, r, 400, "invalid request payload")
		return
	}

	// TODO: data validation of req sturct & saving

}

func GetAllMenu(w http.ResponseWriter, r *http.Request) {
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
}

func PutMenu(w http.ResponseWriter, r *http.Request) {
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
}
