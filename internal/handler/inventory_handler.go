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
	var reqBody []*models.InventoryItem

	data, err := io.ReadAll(r.Body)
	if err != nil {
		handler_utils.ErrorJSONResponse(w, r, http.StatusBadRequest, "error reading request body")
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(data, &reqBody); err != nil {
		var singleItem models.InventoryItem
		if err := json.Unmarshal(data, &singleItem); err != nil {
			handler_utils.ErrorJSONResponse(w, r, http.StatusBadRequest, "invalid request payload")
			return
		}
		reqBody = []*models.InventoryItem{&singleItem}
	}

	for _, item := range reqBody {
		if err := item.IsValid(); err != nil {
			handler_utils.ErrorJSONResponse(w, r, http.StatusBadRequest, fmt.Sprintf("%s: %v", item.Name, err))
			return
		}
	}

	if err := service.SaveInventoryItem(reqBody); err != nil {
		handler_utils.ErrorJSONResponse(w, r, http.StatusBadRequest, fmt.Sprint(err))
		return
	}

	handler_utils.SuccessJSONResponse(w, r, http.StatusOK, "successfully updated the inventory")
}

func GetAllInventory(w http.ResponseWriter, r *http.Request) {
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
}

func PutInventory(w http.ResponseWriter, r *http.Request) {
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
}
