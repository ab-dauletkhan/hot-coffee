package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/handler/handler_utils"
	"github.com/ab-dauletkhan/hot-coffee/internal/service"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func PostInventory(w http.ResponseWriter, r *http.Request) {
	req := models.InventoryItem{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler_utils.JSONResponse(w, r, 400, "invalid request payload", "error")
		return
	}

	if err := req.IsValid(); err != nil {
		handler_utils.JSONResponse(w, r, 400, fmt.Sprint(err), "error")
		return
	}

	if err := service.SaveInventoryItem(req); err != nil {
		handler_utils.JSONResponse(w, r, 400, fmt.Sprint(err), "error")
		return
	}

	handler_utils.JSONResponse(w, r, 200, "successfully updated the inventory", "success")
}

func GetAllInventory(w http.ResponseWriter, r *http.Request) {
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
}

func PutInventory(w http.ResponseWriter, r *http.Request) {
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
}
