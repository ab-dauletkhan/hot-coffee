package handler

import (
	"encoding/json"
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
	req := models.MenuItem{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorJSONResponse(w, r, 400, "invalid request payload")
		return
	}

	// TODO: data validation of req sturct & saving
}

func (h MenuHandler) GetAllMenu(w http.ResponseWriter, r *http.Request) {
}

func (h MenuHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
}

func (h MenuHandler) PutMenuItem(w http.ResponseWriter, r *http.Request) {
}

func (h MenuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
}
