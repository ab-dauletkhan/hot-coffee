package repository

import (
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id string) (*models.Order, error)
	GetAll() ([]*models.Order, error)
	Update(order *models.Order) error
	Delete(id string) error
}

// OrderRepository manages order data
type orderRepository struct {
	storage *JSONStorage
	log     *slog.Logger
}

// NewOrderRepository initializes an OrderRepository with storage and logging
func NewOrderRepository(storage *JSONStorage, log *slog.Logger) *orderRepository {
	return &orderRepository{
		storage: storage,
		log:     log,
	}
}

func (r *orderRepository) Create(order *models.Order) error {
	// 1. Generate ID if needed
	// 2. Read existing data
	// 3. Append new order
	// 4. Write back to file
	// 5. Log operation
	return nil
}

func (r *orderRepository) GetByID(id string) (*models.Order, error) {
	return nil, nil
}
func (r *orderRepository) GetAll() ([]*models.Order, error) {
	return nil, nil
}
func (r *orderRepository) Update(order *models.Order) error {
	return nil
}
func (r *orderRepository) Delete(id string) error {
	return nil
}
