package handler

import (
	"encoding/json"
	"fmt"
	"io"
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
	h.log.Info("CreateOrder called")

	// Read request body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error(fmt.Sprintf("error reading request body: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	// Unmarshal request body into Order struct
	var order models.Order
	if err = json.Unmarshal(data, &order); err != nil {
		h.log.Error(fmt.Sprintf("error unmarshalling order: %v", err))
		writeError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// Validate order
	if err := order.IsValid(); err != nil {
		h.log.Error(fmt.Sprintf("invalid order: %v", err))
		writeError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// Send order to order service
	if err := h.orderService.CreateOrder(&order); err != nil {
		h.log.Error(fmt.Sprintf("error creating order: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("order created: %v", order))
	writeJSON(w, http.StatusCreated, order)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	h.log.Info("GetAllOrders called")

	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		h.log.Error(fmt.Sprintf("error getting orders: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("orders retrieved: %v", orders))
	writeJSON(w, http.StatusOK, orders)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	h.log.Info("GetOrder called")

	id := r.PathValue("id")
	order, err := h.orderService.GetOrder(id)
	if err != nil {
		h.log.Error(fmt.Sprintf("error getting order: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if order == nil {
		h.log.Error(fmt.Sprintf("order not found: %s", id))
		writeError(w, http.StatusNotFound, "Order not found")
		return
	}

	h.log.Info(fmt.Sprintf("order retrieved: %v", order))
	writeJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) PutOrder(w http.ResponseWriter, r *http.Request) {
	h.log.Info("PutOrder called")

	id := r.PathValue("id")

	// Read request body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error(fmt.Sprintf("error reading request body: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	// Unmarshal request body into Order struct
	var order models.Order
	if err = json.Unmarshal(data, &order); err != nil {
		h.log.Error(fmt.Sprintf("error unmarshalling order: %v", err))
		writeError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// Validate order
	if err := order.IsValid(); err != nil {
		h.log.Error(fmt.Sprintf("invalid order: %v", err))
		writeError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// Send order to order service
	if err := h.orderService.UpdateOrder(id, &order); err != nil {
		h.log.Error(fmt.Sprintf("error updating order: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("order updated: %v", order))
	writeJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	h.log.Info("DeleteOrder called")

	id := r.PathValue("id")
	if err := h.orderService.DeleteOrder(id); err != nil {
		h.log.Error(fmt.Sprintf("error deleting order: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("order deleted: %s", id))
	writeJSON(w, http.StatusOK, response{Data: fmt.Sprintf("order deleted: %s", id)})
}

func (h *OrderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	h.log.Info("CloseOrder called")

	id := r.PathValue("id")
	if err := h.orderService.CloseOrder(id); err != nil {
		h.log.Error(fmt.Sprintf("error closing order: %v", err))
		writeError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.log.Info(fmt.Sprintf("order closed: %s", id))
	writeJSON(w, http.StatusOK, response{Data: fmt.Sprintf("order closed: %s", id)})
}
