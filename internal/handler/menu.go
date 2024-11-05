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

// MenuHandler handles HTTP requests for menu items
type MenuHandler struct {
	menuService service.MenuService
	log         *slog.Logger
}

func NewMenuHandler(menuService service.MenuService, log *slog.Logger) *MenuHandler {
	return &MenuHandler{
		menuService: menuService,
		log:         log,
	}
}

func (h MenuHandler) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	h.log.Info("CreateMenuItem called")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error(fmt.Sprintf("error reading request body: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer r.Body.Close()

	// Try to unmarshal into a single item first
	var singleItem models.MenuItem
	if err := json.Unmarshal(data, &singleItem); err == nil {
		// If successful, handle single item addition
		h.handleSingleMenuItem(singleItem, w)
		return
	}

	// Otherwise, try to unmarshal into an array of items
	var items []models.MenuItem
	if err := json.Unmarshal(data, &items); err != nil {
		h.log.Error(fmt.Sprintf("error unmarshalling request body: %v", err))
		writeError(w, http.StatusBadRequest, "Invalid format: expected single or multiple menu items")
		return
	}

	// Handle multiple items
	h.handleMultipleMenuItems(items, w)
}

func (h MenuHandler) handleSingleMenuItem(item models.MenuItem, w http.ResponseWriter) {
	h.log.Info("handleSingleMenuItem called")

	if err := item.IsValid(); err != nil {
		h.log.Error(fmt.Sprintf("error validating menu item: %v", err))
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.menuService.CreateMenuItem(&item); err != nil {
		if errors.Is(err, service.ErrInventoryItemExists) {
			h.log.Error(fmt.Sprintf("menu item already exists: %v", item))
			writeError(w, http.StatusConflict, "Menu item already exists")
			return
		}
		h.log.Error(fmt.Sprintf("error adding menu item: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("menu item added: %v", item))
	w.WriteHeader(http.StatusCreated)
}

func (h MenuHandler) handleMultipleMenuItems(items []models.MenuItem, w http.ResponseWriter) {
	h.log.Info("handleMultipleMenuItems called")

	for _, item := range items {
		if err := item.IsValid(); err != nil {
			h.log.Error(fmt.Sprintf("error validating menu item: %v", err))
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	if errItem, err := h.menuService.CreateMenuItems(&items); err != nil {
		if errors.Is(err, service.ErrInventoryItemExists) {
			h.log.Error(fmt.Sprintf("some menu item already exists: %v", errItem))
			writeError(w, http.StatusConflict, fmt.Sprintf("%s already exists", errItem.Name))
			return
		}
		h.log.Error(fmt.Sprintf("error adding menu items: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("menu items added: %v", items))
	w.WriteHeader(http.StatusCreated)
}

func (h MenuHandler) GetAllMenu(w http.ResponseWriter, r *http.Request) {
	h.log.Info("GetAllMenu called")

	items, err := h.menuService.GetAllMenuItems()
	if err != nil {
		h.log.Error(fmt.Sprintf("error getting menu items: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("menu items retrieved: %v", items))
	writeJSON(w, http.StatusOK, items)
}

func (h MenuHandler) GetAvailableMenuItems(w http.ResponseWriter, r *http.Request) {
	h.log.Info("GetAvailableMenuItems called")

	items, err := h.menuService.GetAvailableMenuItems()
	if err != nil {
		h.log.Error(fmt.Sprintf("error getting menu items: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("menu items retrieved: %v", items))
	writeJSON(w, http.StatusOK, items)
}

func (h MenuHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
	h.log.Info("GetMenuItem called")

	id := r.PathValue("id")
	item, err := h.menuService.GetMenuItem(id)
	if err != nil {
		h.log.Error(fmt.Sprintf("error getting menu item: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if item == nil {
		h.log.Error(fmt.Sprintf("menu item not found: %s", id))
		writeError(w, http.StatusNotFound, "Menu item not found")
		return
	}

	h.log.Info(fmt.Sprintf("menu item retrieved: %v", item))
	writeJSON(w, http.StatusOK, item)
}

func (h MenuHandler) PutMenuItem(w http.ResponseWriter, r *http.Request) {
	h.log.Info("PutMenuItem called")

	id := r.PathValue("id")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error(fmt.Sprintf("error reading request body: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	var item models.MenuItem
	if err := json.Unmarshal(data, &item); err != nil {
		h.log.Error(fmt.Sprintf("error unmarshalling request body: %v", err))
		writeError(w, http.StatusBadRequest, "Invalid format: expected menu item")
		return
	}

	if err := item.IsValid(); err != nil {
		h.log.Error(fmt.Sprintf("error validating menu item: %v", err))
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if item.ID != id {
		h.log.Error(fmt.Sprintf("id mismatch: %s != %s", item.ID, id))
		writeError(w, http.StatusBadRequest, "ID mismatch")
		return
	}

	if err := h.menuService.UpdateMenuItem(id, &item); err != nil {
		if errors.Is(err, service.ErrMenuItemNotFound) {
			h.log.Error(fmt.Sprintf("menu item not found: %s", id))
			writeError(w, http.StatusNotFound, "Menu item not found")
			return
		}

		h.log.Error(fmt.Sprintf("error updating menu item: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func (h MenuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	h.log.Info("DeleteMenuItem called")

	id := r.PathValue("id")
	err := h.menuService.DeleteMenuItem(id)
	if err != nil {
		if errors.Is(err, service.ErrMenuItemNotFound) {
			h.log.Error(fmt.Sprintf("menu item not found: %s", id))
			writeError(w, http.StatusNotFound, "Menu item not found")
			return
		}

		h.log.Error(fmt.Sprintf("error deleting menu item: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("menu item deleted: %s", id))
	writeJSON(w, http.StatusNoContent, nil)
}
