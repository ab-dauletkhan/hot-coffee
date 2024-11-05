package repository

import (
	"fmt"
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

type InventoryRepository interface {
	Create(item *models.InventoryItem) error
	GetByID(id string) (*models.InventoryItem, error)
	GetAll() (*[]models.InventoryItem, error)
	Update(item *models.InventoryItem) error
	Delete(id string) error
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

// loadItems is a helper function to retrieve items from storage
func (r *inventoryRepository) loadItems() (*[]models.InventoryItem, error) {
	var items []models.InventoryItem
	if err := r.storage.Retrieve(&items); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrStorageOperation, err)
	}
	return &items, nil
}

// saveItems is a helper function to save items to storage
func (r *inventoryRepository) saveItems(items *[]models.InventoryItem) error {
	if err := r.storage.Save(items); err != nil {
		return fmt.Errorf("%w: %v", ErrStorageOperation, err)
	}
	return nil
}

func (r *inventoryRepository) Create(item *models.InventoryItem) error {
	r.log.Info("creating inventory item", "id", item.IngredientID)

	items, err := r.loadItems()
	if err != nil {
		return err
	}

	*items = append(*items, *item)

	if err := r.saveItems(items); err != nil {
		r.log.Error("failed to save new inventory item", "error", err, "id", item.IngredientID)
		return err
	}

	return nil
}

func (r *inventoryRepository) GetByID(id string) (*models.InventoryItem, error) {
	r.log.Info("retrieving inventory item", "id", id)

	items, err := r.loadItems()
	if err != nil {
		return nil, err
	}

	for _, item := range *items {
		if item.IngredientID == id {
			return &item, nil
		}
	}

	return nil, nil
}

func (r *inventoryRepository) GetAll() (*[]models.InventoryItem, error) {
	r.log.Info("retrieving all inventory items")

	items, err := r.loadItems()
	if err != nil {
		r.log.Error("failed to load inventory items", "error", err)
		return nil, err
	}

	return items, nil
}

func (r *inventoryRepository) Update(item *models.InventoryItem) error {
	r.log.Info("updating inventory item", "id", item.IngredientID)

	items, err := r.loadItems()
	if err != nil {
		return err
	}

	for i, existing := range *items {
		if existing.IngredientID == item.IngredientID {
			(*items)[i] = *item
			break
		}
	}

	if err := r.saveItems(items); err != nil {
		r.log.Error("failed to save updated inventory item", "error", err, "id", item.IngredientID)
		return err
	}

	return nil
}

func (r *inventoryRepository) Delete(id string) error {
	r.log.Info("deleting inventory item", "id", id)

	items, err := r.loadItems()
	if err != nil {
		return err
	}

	for i, item := range *items {
		if item.IngredientID == id {
			(*items)[i], (*items)[len(*items)-1] = (*items)[len(*items)-1], (*items)[i]
			*items = (*items)[:len(*items)-1]

			if err := r.saveItems(items); err != nil {
				r.log.Error("failed to save updated inventory items", "error", err)
				return err
			}

			break
		}
	}

	return nil
}
