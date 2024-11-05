package service

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/internal/repository"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

type InventoryService interface {
	CreateInventoryItems(items *[]models.InventoryItem) error
	CreateInventoryItem(item *models.InventoryItem) error
	GetInventoryItem(id string) (*models.InventoryItem, error)
	GetAllInventoryItems() (*[]models.InventoryItem, error)
	UpdateInventoryItem(id string, item *models.InventoryItem) error
	DeleteInventoryItem(id string) error

	CheckIngredients(ingredients []models.MenuItemIngredient, quantity int) (bool, error)
	DeductIngredients(ingredients []models.MenuItemIngredient, quantity int) error
}

var (
	ErrInvalidInput          = errors.New("invalid input parameter")
	ErrInsufficientQuantity  = errors.New("insufficient quantity in inventory")
	ErrInventoryItemExists   = errors.New("inventory item already exists")
	ErrInventoryItemNotFound = errors.New("inventory item not found")
)

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

func (s inventoryService) CreateInventoryItems(items *[]models.InventoryItem) error {
	s.log.Info("creating multiple inventory items", "count", len(*items))

	for i := range *items {
		if err := s.CreateInventoryItem(&(*items)[i]); err != nil {
			return fmt.Errorf("failed to create item at index %d: %w", i, err)
		}
	}
	return nil
}

func (s inventoryService) CreateInventoryItem(item *models.InventoryItem) error {
	s.log.Info("creating inventory item", "id", item.IngredientID)

	existingItem, err := s.inventoryRepo.GetByID(item.IngredientID)
	if err != nil {
		s.log.Error("failed to check existing item", "error", err, "id", item.IngredientID)
		return fmt.Errorf("failed to check existing item: %w", err)
	}

	if existingItem != nil {
		s.log.Info("item already exists", "id", item.IngredientID)
		return ErrInventoryItemExists
	}

	if err := s.inventoryRepo.Create(item); err != nil {
		s.log.Error("failed to create item", "error", err, "id", item.IngredientID)
		return fmt.Errorf("failed to create item: %w", err)
	}
	return nil
}

func (s inventoryService) GetInventoryItem(id string) (*models.InventoryItem, error) {
	s.log.Info("retrieving inventory item", "id", id)

	item, err := s.inventoryRepo.GetByID(id)
	if err != nil {
		s.log.Error("failed to get item", "error", err, "id", id)
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	if item == nil {
		return nil, ErrInventoryItemNotFound
	}

	return item, nil
}

func (s inventoryService) GetAllInventoryItems() (*[]models.InventoryItem, error) {
	s.log.Info("retrieving all inventory items")

	items, err := s.inventoryRepo.GetAll()
	if err != nil {
		s.log.Error("failed to get all items", "error", err)
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}
	return items, nil
}

func (s inventoryService) UpdateInventoryItem(id string, item *models.InventoryItem) error {
	s.log.Info("updating inventory item", "id", id)

	existingItem, err := s.inventoryRepo.GetByID(id)
	if err != nil {
		s.log.Error("failed to check existing item", "error", err, "id", id)
		return fmt.Errorf("failed to check existing item: %w", err)
	}

	if existingItem == nil {
		return ErrInventoryItemNotFound
	}

	if err := s.inventoryRepo.Update(item); err != nil {
		s.log.Error("failed to update item", "error", err, "id", id)
		return fmt.Errorf("failed to update item: %w", err)
	}
	return nil
}

func (s inventoryService) DeleteInventoryItem(id string) error {
	s.log.Info("deleting inventory item", "id", id)

	existingItem, err := s.inventoryRepo.GetByID(id)
	if err != nil {
		s.log.Error("failed to check existing item", "error", err, "id", id)
		return fmt.Errorf("failed to check existing item: %w", err)
	}

	if existingItem == nil {
		return ErrInventoryItemNotFound
	}

	if err := s.inventoryRepo.Delete(id); err != nil {
		s.log.Error("failed to delete item", "error", err, "id", id)
		return fmt.Errorf("failed to delete item: %w", err)
	}
	return nil
}

func (s inventoryService) CheckIngredients(ingredients []models.MenuItemIngredient, quantity int) (bool, error) {
	s.log.Info("checking ingredients availability", "ingredients_count", len(ingredients), "quantity", quantity)

	for _, ingredient := range ingredients {
		item, err := s.inventoryRepo.GetByID(ingredient.IngredientID)
		if err != nil {
			s.log.Error("failed to get item for availability check",
				"error", err, "ingredient_id", ingredient.IngredientID)
			return false, fmt.Errorf("failed to check ingredient %s: %w", ingredient.IngredientID, err)
		}

		if item == nil {
			s.log.Info("ingredient not found", "ingredient_id", ingredient.IngredientID)
			return false, nil
		}

		requiredQuantity := ingredient.Quantity * float64(quantity)
		if item.Quantity < requiredQuantity {
			s.log.Info("insufficient quantity",
				"ingredient_id", ingredient.IngredientID,
				"available", item.Quantity,
				"required", requiredQuantity)
			return false, nil
		}
	}

	return true, nil
}

func (s inventoryService) DeductIngredients(ingredients []models.MenuItemIngredient, quantity int) error {
	s.log.Info("deducting ingredients", "ingredients_count", len(ingredients), "quantity", quantity)

	// Then deduct all ingredients
	for _, ingredient := range ingredients {
		item, err := s.inventoryRepo.GetByID(ingredient.IngredientID)
		if err != nil {
			s.log.Error("failed to get item for deduction",
				"error", err,
				"ingredient_id", ingredient.IngredientID)
			return fmt.Errorf("failed to deduct ingredient %s: %w",
				ingredient.IngredientID, err)
		}

		if item == nil {
			s.log.Info("ingredient not found", "ingredient_id", ingredient.IngredientID)
			return fmt.Errorf("ingredient not found: %s", ingredient.IngredientID)
		}

		item.Quantity -= ingredient.Quantity * float64(quantity)
		if err := s.inventoryRepo.Update(item); err != nil {
			s.log.Error("failed to update item after deduction",
				"error", err,
				"ingredient_id", ingredient.IngredientID)
			return fmt.Errorf("failed to update ingredient %s: %w",
				ingredient.IngredientID, err)
		}
	}

	return nil
}
