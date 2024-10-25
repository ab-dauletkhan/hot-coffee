package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/handler/handler_utils"
	"github.com/ab-dauletkhan/hot-coffee/internal/service"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func PostInventory(w http.ResponseWriter, r *http.Request) {
	req := []models.InventoryItem{}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		handler_utils.ErrorJSONResponse(w, r, 400, fmt.Sprintf("error reading your request body: %v", err))
		return
	}
	if err := json.Unmarshal(data, &req); err != nil {
		temp := models.InventoryItem{}
		if err := json.Unmarshal(data, &temp); err != nil {
			handler_utils.ErrorJSONResponse(w, r, 400, "invalid request payload")
			return
		}
		req = []models.InventoryItem{}
		req = append(req, temp)
	}
	for _, item := range req {
		if err := item.IsValid(); err != nil {
			handler_utils.ErrorJSONResponse(w, r, 400, fmt.Sprintf("%s: %v", item.Name, err))
			return
		}
	}
	if err := service.SaveInventoryItem(&req); err != nil {
		handler_utils.ErrorJSONResponse(w, r, 400, fmt.Sprint(err))
		return
	}

	handler_utils.SuccessJSONResponse(w, r, 200, "successfully updated the inventory")
}

func GetAllInventory(w http.ResponseWriter, r *http.Request) {
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
}

func PutInventory(w http.ResponseWriter, r *http.Request) {
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
}
