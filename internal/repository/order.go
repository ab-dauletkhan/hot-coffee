package repository

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id string) (*models.Order, error)
	GetAll() (*[]models.Order, error)
	Update(order *models.Order) error
	Delete(id string) error
	Close(id string) error
}

// OrderRepository manages order data
type orderRepository struct {
	storage *JSONStorage
	log     *slog.Logger
}

var (
	ErrOrderNotFound  = errors.New("order not found")
	ErrOrderExists    = errors.New("order already exists")
	ErrOrderClosed    = errors.New("order is already closed")
	ErrOrderNotClosed = errors.New("order is not closed")
)

// NewOrderRepository initializes an OrderRepository with storage and logging
func NewOrderRepository(storage *JSONStorage, log *slog.Logger) *orderRepository {
	return &orderRepository{
		storage: storage,
		log:     log,
	}
}

// loadOrders retrieves all orders from storage
func (r *orderRepository) loadOrders() (*[]models.Order, error) {
	var orders []models.Order
	if err := r.storage.Retrieve(&orders); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrStorageOperation, err)
	}
	return &orders, nil
}

// saveOrders saves orders to storage
func (r *orderRepository) saveOrders(orders *[]models.Order) error {
	if err := r.storage.Save(orders); err != nil {
		return fmt.Errorf("%w: %v", ErrStorageOperation, err)
	}
	return nil
}

func (r *orderRepository) Create(order *models.Order) error {
	r.log.Info("creating new order", "order_id", order.ID)

	orders, err := r.loadOrders()
	if err != nil {
		r.log.Error("failed to load orders", "error", err)
		return err
	}

	*orders = append(*orders, *order)

	if err := r.saveOrders(orders); err != nil {
		r.log.Error("failed to save new order", "error", err, "order_id", order.ID)
		return err
	}

	return nil
}

func (r *orderRepository) GetByID(id string) (*models.Order, error) {
	r.log.Info("retrieving order", "order_id", id)

	orders, err := r.loadOrders()
	if err != nil {
		r.log.Error("failed to load orders", "error", err)
		return nil, err
	}

	for _, order := range *orders {
		if order.ID == id {
			orderCopy := order
			return &orderCopy, nil
		}
	}

	return nil, ErrOrderNotFound
}

func (r *orderRepository) GetAll() (*[]models.Order, error) {
	r.log.Info("retrieving all orders")

	orders, err := r.loadOrders()
	if err != nil {
		r.log.Error("failed to load orders", "error", err)
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) Update(order *models.Order) error {
	r.log.Info("updating order", "order_id", order.ID)

	orders, err := r.loadOrders()
	if err != nil {
		r.log.Error("failed to load orders", "error", err)
		return err
	}

	found := false
	for i := range *orders {
		if (*orders)[i].ID == order.ID {
			(*orders)[i] = *order
			found = true
			break
		}
	}

	if !found {
		return ErrOrderNotFound
	}

	if err := r.saveOrders(orders); err != nil {
		r.log.Error("failed to save updated order", "error", err, "order_id", order.ID)
		return err
	}

	return nil
}

func (r *orderRepository) Delete(id string) error {
	r.log.Info("deleting order", "order_id", id)

	orders, err := r.loadOrders()
	if err != nil {
		r.log.Error("failed to load orders", "error", err)
		return err
	}

	for i := range *orders {
		if (*orders)[i].ID == id {
			// Remove the order by swapping with the last element and truncating
			lastIdx := len(*orders) - 1
			(*orders)[i] = (*orders)[lastIdx]
			*orders = (*orders)[:lastIdx]

			if err := r.saveOrders(orders); err != nil {
				r.log.Error("failed to save orders after deletion", "error", err, "order_id", id)
				return err
			}

			return nil
		}
	}

	return ErrOrderNotFound
}

func (r *orderRepository) Close(id string) error {
	r.log.Info("closing order", "order_id", id)

	orders, err := r.loadOrders()
	if err != nil {
		r.log.Error("failed to load orders", "error", err)
		return err
	}

	for i := range *orders {
		if (*orders)[i].ID == id {
			(*orders)[i].Status = "closed"

			if err := r.saveOrders(orders); err != nil {
				r.log.Error("failed to save orders after closing", "error", err, "order_id", id)
				return err
			}

			return nil
		}
	}

	return ErrOrderNotFound
}
