package service

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/internal/repository"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

type InventoryService interface {
	CreateInventoryItems(items []*models.InventoryItem) error
	CreateInventoryItem(item *models.InventoryItem) error
	GetInventoryItem(id string) (*models.InventoryItem, error)
	GetAllInventoryItems() (*[]models.InventoryItem, error)
	UpdateInventoryItem(id string, item *models.InventoryItem) error
	DeleteInventoryItem(id string) error
	CheckIngredientAvailability(ingredients []models.MenuItemIngredient) error
	DeductIngredients(ingredients []models.MenuItemIngredient) error
	IsNotFoundError(err error) bool
}

var ErrInventoryItemNotFound = errors.New("inventory item not found")

// InventoryService handles business logic for inventory items
type inventoryService struct {
	inventoryRepo repository.InventoryRepository
	log           *slog.Logger
}

// NewInventoryService initializes InventoryService with repository and logging
func NewInventoryService(inventoryRepo repository.InventoryRepository, log *slog.Logger) inventoryService {
	return inventoryService{
		inventoryRepo: inventoryRepo,
		log:           log,
	}
}

func (s inventoryService) CreateInventoryItems(items []*models.InventoryItem) error {
	return nil
}

func (s inventoryService) CreateInventoryItem(item *models.InventoryItem) error {
	s.log.Info("CreateInventoryItem called")
	// 1. Check if the inventory item already exists
	existingItem, err := s.inventoryRepo.GetByID(item.IngredientID)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to check existing inventory item: %v", err))
		return fmt.Errorf("failed to check existing inventory item: %w", err)
	}
	if existingItem != nil {
		s.log.Error(fmt.Sprintf("inventory item with ID %s already exists", item.IngredientID))
		return fmt.Errorf("inventory item with ID %s already exists", item.IngredientID)
	}

	// 2. Additional validation or calculations if needed (e.g., setting initial stock status)

	// 3. Call the repository to save the new item
	if err := s.inventoryRepo.Create(item); err != nil {
		s.log.Error(fmt.Sprintf("failed to create inventory item: %v", err))
		return fmt.Errorf("failed to create inventory item: %w", err)
	}
	return nil
}

func (s inventoryService) GetInventoryItem(id string) (*models.InventoryItem, error) {
	s.log.Info("GetInventoryItem called")

	item, err := s.inventoryRepo.GetByID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to get inventory item: %v", err))
		return nil, fmt.Errorf("failed to get inventory item: %w", err)
	}
	return item, nil
}

func (s inventoryService) GetAllInventoryItems() (*[]models.InventoryItem, error) {
	s.log.Info("GetAllInventoryItems called")
	items, err := s.inventoryRepo.GetAll()
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to get all inventory items: %v", err))
		return nil, fmt.Errorf("failed to get all inventory items: %w", err)
	}
	return items, nil
}

func (s inventoryService) UpdateInventoryItem(id string, item *models.InventoryItem) error {
	s.log.Info("UpdateInventoryItem called")

	// 1. Check if the inventory item exists
	existingItem, err := s.inventoryRepo.GetByID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to check existing inventory item: %v", err))
		return fmt.Errorf("failed to check existing inventory item: %w", err)
	}

	if existingItem == nil {
		s.log.Error(fmt.Sprintf("inventory item with ID %s not found", id))
		return ErrInventoryItemNotFound
	}

	// 2. Additional validation or calculations if needed (e.g., setting initial stock status)

	// 3. Call the repository to save the new item
	if err := s.inventoryRepo.Update(item); err != nil {
		s.log.Error(fmt.Sprintf("failed to update inventory item: %v", err))
		return fmt.Errorf("failed to update inventory item: %w", err)
	}
	return nil
}

func (s inventoryService) DeleteInventoryItem(id string) error {
	s.log.Info("DeleteInventoryItem called")

	// 1. Check if the inventory item exists
	existingItem, err := s.inventoryRepo.GetByID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to check existing inventory item: %v", err))
		return fmt.Errorf("failed to check existing inventory item: %w", err)
	}

	if existingItem == nil {
		s.log.Error(fmt.Sprintf("inventory item with ID %s not found", id))
		return ErrInventoryItemNotFound
	}

	// 2. Call the repository to delete the item
	if err := s.inventoryRepo.Delete(id); err != nil {
		s.log.Error(fmt.Sprintf("failed to delete inventory item: %v", err))
		return fmt.Errorf("failed to delete inventory item: %w", err)
	}
	return nil
}

func (s inventoryService) CheckIngredientAvailability(ingredients []models.MenuItemIngredient) error {
	return nil
}

func (s inventoryService) DeductIngredients(ingredients []models.MenuItemIngredient) error {
	return nil
}

func (s inventoryService) IsNotFoundError(err error) bool {
	return errors.Is(err, ErrInventoryItemNotFound)
}

// func SaveInventoryItem(items []*models.InventoryItem) error {
// 	inventoryItems, err := repository.GetJSONInventory()
// 	if err != nil {
// 		return err
// 	}

// 	itemMap := mapInventoryItems(inventoryItems)

// 	for _, item := range items {
// 		if err := updateOrAddItem(item, itemMap); err != nil {
// 			return err
// 		}
// 	}

// 	err = repository.SaveJSONInventoryItem(mapToSlice(itemMap))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func mapInventoryItems(items []*models.InventoryItem) map[string]*models.InventoryItem {
// 	itemMap := make(map[string]*models.InventoryItem)
// 	for _, item := range items {
// 		itemMap[item.IngredientID] = item
// 	}
// 	return itemMap
// }

// func updateOrAddItem(item *models.InventoryItem, itemMap map[string]*models.InventoryItem) error {
// 	if existingItem, exists := itemMap[item.IngredientID]; exists {
// 		if existingItem.Unit != item.Unit {
// 			return fmt.Errorf("%s: invalid ingredient unit - have (%s), want (%s)",
// 				item.Name, item.Unit, existingItem.Unit)
// 		}
// 		existingItem.Quantity += item.Quantity
// 	} else {
// 		itemMap[item.IngredientID] = item
// 	}
// 	return nil
// }

// func mapToSlice(itemMap map[string]*models.InventoryItem) []*models.InventoryItem {
// 	items := make([]*models.InventoryItem, 0, len(itemMap))
// 	for _, item := range itemMap {
// 		items = append(items, item)
// 	}
// 	return items
// }
