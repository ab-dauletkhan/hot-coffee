package repository

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

type InventoryRepository interface {
	Create(item *models.InventoryItem) error
	GetByID(id string) (*models.InventoryItem, error)
	GetAll() ([]*models.InventoryItem, error)
	Update(item *models.InventoryItem) error
	Delete(id string) error
	UpdateQuantity(id string, delta float64) error
}

// InventoryRepository manages inventory data
type inventoryRepository struct {
	storage *JSONStorage
	log     *slog.Logger
}

// NewInventoryRepository initializes an InventoryRepository with storage and logging
func NewInventoryRepository(storage *JSONStorage, log *slog.Logger) *inventoryRepository {
	return &inventoryRepository{
		storage: storage,
		log:     log,
	}
}

func (r *inventoryRepository) Create(item *models.InventoryItem) error {
	// 1. 

	return nil
}
func (r *inventoryRepository) GetByID(id string) (*models.InventoryItem, error) {
	return nil, nil
}
func (r *inventoryRepository) GetAll() ([]*models.InventoryItem, error) {
	return nil, nil
}
func (r *inventoryRepository) Update(item *models.InventoryItem) error {
	return nil
}
func (r *inventoryRepository) Delete(id string) error {
	return nil
}
func (r *inventoryRepository) UpdateQuantity(id string, delta float64) error {
	return nil
}

const inventoryFilePath = "data/inventory.json"

func GetJSONInventory() ([]*models.InventoryItem, error) {
	data, err := os.ReadFile(inventoryFilePath)
	if err != nil {
		slog.Debug(fmt.Sprintf("error reading %s: %v", inventoryFilePath, err))
		return nil, fmt.Errorf("failed to read inventory file: %w", err)
	}

	var inventoryItems []*models.InventoryItem
	if err := json.Unmarshal(data, &inventoryItems); err != nil {
		slog.Debug(fmt.Sprintf("error unmarshalling inventory data: %v", err))
		return nil, fmt.Errorf("failed to parse inventory data: %w", err)
	}

	return inventoryItems, nil
}

func SaveJSONInventoryItem(items []*models.InventoryItem) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		slog.Debug(fmt.Sprintf("error marshalling inventory items: %v", err))
		return fmt.Errorf("failed to save inventory data: %w", err)
	}

	if err := os.WriteFile(inventoryFilePath, data, filePerm); err != nil {
		slog.Debug(fmt.Sprintf("error writing to %s: %v", inventoryFilePath, err))
		return fmt.Errorf("failed to write inventory file: %w", err)
	}

	return nil
}
