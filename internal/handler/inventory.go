package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/repository"
	"github.com/ab-dauletkhan/hot-coffee/internal/service"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

// InventoryHandler handles HTTP requests for inventory
type InventoryHandler struct {
	inventoryService service.InventoryService
	log              *slog.Logger
}

func NewInventoryHandler(inventoryService service.InventoryService, log *slog.Logger) *InventoryHandler {
	return &InventoryHandler{
		inventoryService: inventoryService,
		log:              log,
	}
}

func (h InventoryHandler) AddInventoryItem(w http.ResponseWriter, r *http.Request) {
	var reqBody []*models.InventoryItem

	data, err := io.ReadAll(r.Body)
	if err != nil {
		ErrorJSONResponse(w, r, http.StatusBadRequest, "error reading request body")
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(data, &reqBody); err != nil {
		fmt.Println(reqBody)
		var singleItem models.InventoryItem
		if err := json.Unmarshal(data, &singleItem); err != nil {
			fmt.Println(singleItem)
			ErrorJSONResponse(w, r, http.StatusBadRequest, "invalid request payload")
			return
		}
		reqBody = []*models.InventoryItem{&singleItem}
	}

	for _, item := range reqBody {
		if err := item.IsValid(); err != nil {
			ErrorJSONResponse(w, r, http.StatusBadRequest, fmt.Sprintf("%s: %v", item.Name, err))
			return
		}
	}

	if err := service.SaveInventoryItem(reqBody); err != nil {
		ErrorJSONResponse(w, r, http.StatusBadRequest, fmt.Sprint(err))
		return
	}

	SuccessJSONResponse(w, r, http.StatusOK, "successfully updated the inventory")
}

func (h InventoryHandler) GetAllInventory(w http.ResponseWriter, r *http.Request) {
	JSONItems, err := repository.GetJSONInventory()
	if err != nil {
		ErrorJSONResponse(w, r, 400, fmt.Sprint(err))
		return
	}

	CustomJSONResponse(w, r, 200, "successful get reponse", nil, JSONItems, slog.LevelInfo)
}

func (h InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
}

func (h InventoryHandler) PutInventory(w http.ResponseWriter, r *http.Request) {
}

func (h InventoryHandler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
}
