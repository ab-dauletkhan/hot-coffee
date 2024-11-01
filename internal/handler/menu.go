package handler

import (
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/service"
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
}

func (h MenuHandler) GetAllMenu(w http.ResponseWriter, r *http.Request) {
}

func (h MenuHandler) GetMenuItem(w http.ResponseWriter, r *http.Request) {
}

func (h MenuHandler) PutMenuItem(w http.ResponseWriter, r *http.Request) {
}

func (h MenuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
}
