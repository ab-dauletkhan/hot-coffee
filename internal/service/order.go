package service

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/ab-dauletkhan/hot-coffee/internal/repository"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

type OrderService interface {
	CreateOrder(order *models.Order) (models.Order, error)
	GetOrder(id string) (*models.Order, error)
	GetAllOrders() (*[]models.Order, error)
	UpdateOrder(id string, order *models.Order) error
	DeleteOrder(id string) error
	CloseOrder(id string) error

	NewOrderID(name string) string
}

// OrderService handles business logic for orders
type orderService struct {
	orderRepo        repository.OrderRepository
	menuService      MenuService
	inventoryService InventoryService
	log              *slog.Logger
}

var (
	ErrOrderNotFound  = errors.New("order not found")
	ErrOrderExists    = errors.New("order already exists")
	ErrOrderClosed    = errors.New("order is already closed")
	ErrOrderNotClosed = errors.New("order is not closed")
)

// NewOrderService initializes OrderService with repositories and logging
func NewOrderService(orderRepo repository.OrderRepository, menuService MenuService, inventoryService InventoryService, log *slog.Logger) orderService {
	return orderService{
		orderRepo:        orderRepo,
		menuService:      menuService,
		inventoryService: inventoryService,
		log:              log,
	}
}

func (r orderService) CreateOrder(order *models.Order) (models.Order, error) {
	r.log.Info("CreateOrder called")

	for _, item := range order.Items {
		ok, err := r.menuService.IsMenuAvailable(item.ProductID, item.Quantity)
		if err != nil {
			if errors.Is(err, ErrMenuItemNotFound) {
				return models.Order{}, fmt.Errorf("product %s not found", item.ProductID)
			}
			return models.Order{}, err
		}

		if !ok {
			return models.Order{}, fmt.Errorf("product %s is not available", item.ProductID)
		}
	}

	for _, item := range order.Items {
		err := r.menuService.PrepareMenu(item.ProductID, item.Quantity)
		if err != nil {
			return models.Order{}, err
		}
	}

	customer_name := strings.ReplaceAll(strings.ToLower(order.CustomerName), " ", "_")
	order.ID = r.NewOrderID(customer_name)
	order.Status = models.StatusPending
	order.CreatedAt = time.Now().Format(time.RFC3339)

	err := r.orderRepo.Create(order)
	if err != nil {
		return models.Order{}, err
	}

	return *order, nil
}

func (r orderService) CloseOrder(id string) error {
	r.log.Info("CloseOrder called")

	order, err := r.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	if order == nil {
		return ErrOrderNotFound
	}

	// for _, item := range order.Items {
	// 	err := r.menuService.PrepareMenu(item.ProductID, item.Quantity)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	err = r.orderRepo.Close(id)
	if err != nil {
		return err
	}
	return nil
}

func (r orderService) GetOrder(id string) (*models.Order, error) {
	r.log.Info("GetOrder called")

	order, err := r.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r orderService) GetAllOrders() (*[]models.Order, error) {
	r.log.Info("GetAllOrders called")

	orders, err := r.orderRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r orderService) UpdateOrder(id string, order *models.Order) error {
	r.log.Info("UpdateOrder called")

	existing, err := r.GetOrder(id)
	if err != nil {
		return err
	}

	if existing == nil {
		return ErrOrderNotFound
	}

	err = r.orderRepo.Update(order)
	if err != nil {
		return err
	}

	return nil
}

func (r orderService) DeleteOrder(id string) error {
	r.log.Info("DeleteOrder called")

	err := r.orderRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (r orderService) NewOrderID(name string) string {
	return fmt.Sprintf("order-%s-%d", name, time.Now().Unix())
}
