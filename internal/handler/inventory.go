package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

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

func (h InventoryHandler) AddInventoryItems(w http.ResponseWriter, r *http.Request) {
	h.log.Info("AddInventoryItems called")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error(fmt.Sprintf("error reading request body: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer r.Body.Close()

	// Try to unmarshal into a single item first
	var singleItem models.InventoryItem
	if err := json.Unmarshal(data, &singleItem); err == nil {
		// If successful, handle single item addition
		h.handleSingleInventoryItem(singleItem, w)
		return
	}

	// Otherwise, try to unmarshal into an array of items
	var items []models.InventoryItem
	if err := json.Unmarshal(data, &items); err != nil {
		h.log.Error(fmt.Sprintf("error unmarshalling request body: %v", err))
		writeError(w, http.StatusBadRequest, "Invalid format: expected single or multiple inventory items")
		return
	}

	// Handle multiple items
	h.handleMultipleInventoryItems(items, w)
}

// Private helper to handle single inventory item addition
func (h InventoryHandler) handleSingleInventoryItem(item models.InventoryItem, w http.ResponseWriter) {
	h.log.Info("handling single inventory item")

	if err := item.IsValid(); err != nil {
		h.log.Error(fmt.Sprintf("invalid inventory item: %v", err))
		writeError(w, http.StatusBadRequest, fmt.Sprintf("%s: %v", item.Name, err))
		return
	}

	if err := h.inventoryService.CreateInventoryItem(&item); err != nil {
		h.log.Error(fmt.Sprintf("error creating single inventory item: %v", err))
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.log.Info("successfully created single inventory item")
	writeJSON(w, http.StatusCreated, item)
}

// Private helper to handle multiple inventory items addition
func (h InventoryHandler) handleMultipleInventoryItems(items []models.InventoryItem, w http.ResponseWriter) {
	h.log.Info("handling multiple inventory items")

	for _, item := range items {
		if err := item.IsValid(); err != nil {
			h.log.Error(fmt.Sprintf("invalid inventory item: %v", err))
			writeError(w, http.StatusBadRequest, fmt.Sprintf("Item %s: %v", item.Name, err))
			return
		}
	}

	if err := h.inventoryService.CreateInventoryItems(&items); err != nil {
		h.log.Error(fmt.Sprintf("error creating multiple inventory items: %v", err))
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.log.Info("successfully created multiple inventory items")
	writeJSON(w, http.StatusCreated, items)
}

func (h InventoryHandler) GetAllInventory(w http.ResponseWriter, r *http.Request) {
	h.log.Info("GetAllInventory called")

	items, err := h.inventoryService.GetAllInventoryItems()
	if err != nil {
		h.log.Error(fmt.Sprintf("error getting inventory items: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	h.log.Info(fmt.Sprintf("got all inventory items: %v", items))
	writeJSON(w, http.StatusOK, items)
}

func (h InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	h.log.Info("GetInventory called")

	id := r.PathValue("id")

	item, err := h.inventoryService.GetInventoryItem(id)
	if err != nil {
		if errors.Is(err, service.ErrInventoryItemNotFound) {
			h.log.Error(fmt.Sprintf("item not found: %s", id))
			writeError(w, http.StatusNotFound, fmt.Sprintf("item not found: %s", id))
			return
		}

		h.log.Error(fmt.Sprintf("error getting inventory item: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("got inventory item: %v", item))
	writeJSON(w, http.StatusOK, item)
}

func (h InventoryHandler) PutInventory(w http.ResponseWriter, r *http.Request) {
	h.log.Info("PutInventory called")

	id := r.PathValue("id")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error(fmt.Sprintf("error reading request body: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer r.Body.Close()

	var item models.InventoryItem
	if err := json.Unmarshal(data, &item); err != nil {
		h.log.Error(fmt.Sprintf("error unmarshalling request body: %v", err))
		writeError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := item.IsValid(); err != nil {
		h.log.Error(fmt.Sprintf("invalid item: %v", err))
		writeError(w, http.StatusBadRequest, fmt.Sprintf("%s: %v", item.Name, err))
		return
	}

	if id != item.IngredientID {
		h.log.Error(fmt.Sprintf("ID mismatch: have (%s), want (%s)", item.IngredientID, id))
		writeError(w, http.StatusBadRequest, "ID mismatch in request body and URL")
		return
	}

	err = h.inventoryService.UpdateInventoryItem(id, &item)
	if err != nil {
		if errors.Is(err, service.ErrInventoryItemNotFound) {
			h.log.Error(fmt.Sprintf("item not found: %s", id))
			writeError(w, http.StatusNotFound, fmt.Sprintf("item not found: %s", id))
			return
		}

		h.log.Error(fmt.Sprintf("error updating inventory item: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("updated inventory item: %v", item))
	writeJSON(w, http.StatusOK, item)
}

func (h InventoryHandler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
	h.log.Info("DeleteInventory called")

	id := r.PathValue("id")

	err := h.inventoryService.DeleteInventoryItem(id)
	if err != nil {
		if errors.Is(err, service.ErrInventoryItemNotFound) {
			h.log.Error(fmt.Sprintf("item not found: %s", id))
			writeError(w, http.StatusNotFound, fmt.Sprintf("item not found: %s", id))
			return
		}

		h.log.Error(fmt.Sprintf("error deleting inventory item: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("deleted inventory item: %s", id))
	writeJSON(w, http.StatusNoContent, nil)
}
