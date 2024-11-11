package service

import (
	"errors"
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/internal/repository"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

type MenuService interface {
	CreateMenuItems(items *[]models.MenuItem) (*models.MenuItem, error)
	CreateMenuItem(item *models.MenuItem) error
	GetMenuItem(id string) (*models.MenuItem, error)
	GetAllMenuItems() (*[]models.MenuItem, error)
	GetAvailableMenuItems() (*[]models.MenuItem, error)
	UpdateMenuItem(id string, item *models.MenuItem) error
	DeleteMenuItem(id string) error

	IsMenuAvailable(id string, quantity int) (bool, error)
	PrepareMenu(id string, quantity int) error

	GetPriceByID(id string) (float64, error)
}

var (
	ErrMenuItemAlreadyExists = errors.New("menu item already exists")
	ErrMenuItemNotFound      = errors.New("menu item not found")
)

// MenuService handles business logic for menu items
type menuService struct {
	menuRepo         repository.MenuRepository
	inventoryService InventoryService
	log              *slog.Logger
}

// NewMenuService initializes MenuService with repository and logging
func NewMenuService(menuRepo repository.MenuRepository, inventoryService InventoryService, log *slog.Logger) menuService {
	return menuService{
		menuRepo:         menuRepo,
		inventoryService: inventoryService,
		log:              log,
	}
}

func (s menuService) CreateMenuItems(items *[]models.MenuItem) (*models.MenuItem, error) {
	s.log.Info("CreateMenuItems called")

	for _, item := range *items {
		if err := s.CreateMenuItem(&item); err != nil {
			return &item, err
		}
	}
	return nil, nil
}

func (s menuService) CreateMenuItem(item *models.MenuItem) error {
	s.log.Info("CreateMenuItem called")

	// 1. Check if item already exists
	existingItem, err := s.menuRepo.GetByID(item.ID)
	if err != nil {
		s.log.Error("failed to check existing menu item")
		return err
	}

	// 2. If item exists, return an error
	if existingItem != nil {
		s.log.Error("menu item already exists")
		return ErrMenuItemAlreadyExists
	}

	// 3. Create the item
	if err := s.menuRepo.Create(item); err != nil {
		s.log.Error("failed to create menu item")
		return err
	}

	return nil
}

func (s menuService) GetMenuItem(id string) (*models.MenuItem, error) {
	s.log.Info("GetMenuItem called")

	item, err := s.menuRepo.GetByID(id)
	if err != nil {
		s.log.Error("failed to get menu item")
		return nil, err
	}

	return item, nil
}

func (s menuService) GetAllMenuItems() (*[]models.MenuItem, error) {
	s.log.Info("GetAllMenuItems called")

	items, err := s.menuRepo.GetAll()
	if err != nil {
		s.log.Error("failed to get all menu items")
		return nil, err
	}

	return items, nil
}

func (s menuService) GetAvailableMenuItems() (*[]models.MenuItem, error) {
	s.log.Info("GetAvailableMenuItems called")

	return nil, nil
}

func (s menuService) UpdateMenuItem(id string, item *models.MenuItem) error {
	s.log.Info("UpdateMenuItem called")

	// 1. Check if item exists
	existingItem, err := s.menuRepo.GetByID(id)
	if err != nil {
		s.log.Error("failed to check existing menu item")
		return err
	}

	// 2. If item does not exist, return an error
	if existingItem == nil {
		s.log.Error("menu item not found")
		return ErrMenuItemNotFound
	}

	// 3. Update the item
	if err := s.menuRepo.Update(item); err != nil {
		s.log.Error("failed to update menu item")
		return err
	}

	return nil
}

func (s menuService) DeleteMenuItem(id string) error {
	s.log.Info("DeleteMenuItem called")

	// 1. Check if item exists
	existingItem, err := s.menuRepo.GetByID(id)
	if err != nil {
		s.log.Error("failed to check existing menu item")
		return err
	}

	// 2. If item does not exist, return an error
	if existingItem == nil {
		s.log.Error("menu item not found")
		return ErrMenuItemNotFound
	}

	// 3. Delete the item
	if err := s.menuRepo.Delete(id); err != nil {
		s.log.Error("failed to delete menu item")
		return err
	}

	return nil
}

func (s menuService) IsMenuAvailable(id string, quantity int) (bool, error) {
	s.log.Info("IsMenuAvailable called")

	item, err := s.menuRepo.GetByID(id)
	if err != nil {
		s.log.Error("failed to get menu item")
		return false, err
	}

	if item == nil {
		s.log.Error("menu item not found")
		return false, ErrMenuItemNotFound
	}

	ok, err := s.inventoryService.CheckIngredients(item.Ingredients, quantity)
	if err != nil {
		s.log.Error("failed to check ingredients availability")
		return false, err
	}

	if !ok {
		s.log.Error("ingredients not available")
		return false, nil
	}

	return true, nil
}

func (s menuService) PrepareMenu(id string, quantity int) error {
	s.log.Info("PrepareMenu called")

	item, err := s.menuRepo.GetByID(id)
	if err != nil {
		s.log.Error("failed to get menu item")
		return err
	}

	if item == nil {
		s.log.Error("menu item not found")
		return ErrMenuItemNotFound
	}

	if err := s.inventoryService.DeductIngredients(item.Ingredients, quantity); err != nil {
		s.log.Error("failed to use ingredients")
		return err
	}

	return nil
}

func (s menuService) GetPriceByID(id string) (float64, error) {
	menuItem, err := s.menuRepo.GetByID(id)
	if err != nil {
		return 0, err
	}
	return menuItem.Price, nil
}
