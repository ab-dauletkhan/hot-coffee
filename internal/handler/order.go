package handler

import (
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/service"
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
	log *slog.Logger,
) *OrderHandler {
	return &OrderHandler{
		orderService:     orderService,
		menuService:      menuService,
		inventoryService: inventoryService,
		log:              log,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) PutOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
}
