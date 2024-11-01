package service

import (
	"fmt"
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/internal/repository"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

type InventoryService interface {
	CreateInventoryItem(item *models.InventoryItem) error
	GetInventoryItem(id string) (*models.InventoryItem, error)
	GetAllInventoryItems() ([]*models.InventoryItem, error)
	UpdateInventoryItem(item *models.InventoryItem) error
	DeleteInventoryItem(id string) error
	CheckIngredientAvailability(ingredients []models.MenuItemIngredient) error
	DeductIngredients(ingredients []models.MenuItemIngredient) error
}

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

func (s inventoryService) CreateInventoryItem(item *models.InventoryItem) error {
	return nil
}
func (s inventoryService) GetInventoryItem(id string) (*models.InventoryItem, error) {
	return nil, nil
}
func (s inventoryService) GetAllInventoryItems() ([]*models.InventoryItem, error) {
	return nil, nil
}
func (s inventoryService) UpdateInventoryItem(item *models.InventoryItem) error {
	return nil
}
func (s inventoryService) DeleteInventoryItem(id string) error {
	return nil
}
func (s inventoryService) CheckIngredientAvailability(ingredients []models.MenuItemIngredient) error {
	return nil
}
func (s inventoryService) DeductIngredients(ingredients []models.MenuItemIngredient) error {
	return nil
}

func SaveInventoryItem(items []*models.InventoryItem) error {
	inventoryItems, err := repository.GetJSONInventory()
	if err != nil {
		return err
	}

	itemMap := mapInventoryItems(inventoryItems)

	for _, item := range items {
		if err := updateOrAddItem(item, itemMap); err != nil {
			return err
		}
	}

	err = repository.SaveJSONInventoryItem(mapToSlice(itemMap))
	if err != nil {
		return err
	}

	return nil
}

func mapInventoryItems(items []*models.InventoryItem) map[string]*models.InventoryItem {
	itemMap := make(map[string]*models.InventoryItem)
	for _, item := range items {
		itemMap[item.IngredientID] = item
	}
	return itemMap
}

func updateOrAddItem(item *models.InventoryItem, itemMap map[string]*models.InventoryItem) error {
	if existingItem, exists := itemMap[item.IngredientID]; exists {
		if existingItem.Unit != item.Unit {
			return fmt.Errorf("%s: invalid ingredient unit - have (%s), want (%s)",
				item.Name, item.Unit, existingItem.Unit)
		}
		existingItem.Quantity += item.Quantity
	} else {
		itemMap[item.IngredientID] = item
	}
	return nil
}

func mapToSlice(itemMap map[string]*models.InventoryItem) []*models.InventoryItem {
	items := make([]*models.InventoryItem, 0, len(itemMap))
	for _, item := range itemMap {
		items = append(items, item)
	}
	return items
}
