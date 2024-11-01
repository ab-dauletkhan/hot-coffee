package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/service"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

type OrderHandler struct {
	orderService     service.OrderService
	menuService      service.MenuService
	inventoryService service.InventoryService
	log              *slog.Logger
}

func NewOrderHandler(orderService service.OrderService,
	menuService service.MenuService,
	inventoryService service.InventoryService,
	log *slog.Logger) *OrderHandler {
	return &OrderHandler{
		orderService:     orderService,
		menuService:      menuService,
		inventoryService: inventoryService,
		log:              log,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	req := models.Order{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorJSONResponse(w, r, 400, "invalid request payload")
		return
	}
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
}

func (h *OrderHandler) PutOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
}
