package repository

import (
	"fmt"
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

type MenuRepository interface {
	Create(item *models.MenuItem) error
	GetByID(id string) (*models.MenuItem, error)
	GetAll() (*[]models.MenuItem, error)
	Update(item *models.MenuItem) error
	Delete(id string) error

	GetRequiredIngredients(id string) (*[]models.MenuItemIngredient, error)
}

// MenuRepository manages menu data
type menuRepository struct {
	storage *JSONStorage
	log     *slog.Logger
}

// NewMenuRepository initializes a MenuRepository with storage and logging
func NewMenuRepository(storage *JSONStorage, log *slog.Logger) *menuRepository {
	return &menuRepository{
		storage: storage,
		log:     log,
	}
}

// loadItems is a helper function to retrieve items from storage
func (r *menuRepository) loadItems() (*[]models.MenuItem, error) {
	var items []models.MenuItem
	if err := r.storage.Retrieve(&items); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrStorageOperation, err)
	}
	return &items, nil
}

// saveItems is a helper function to save items to storage
func (r *menuRepository) saveItems(items *[]models.MenuItem) error {
	if err := r.storage.Save(items); err != nil {
		return fmt.Errorf("%w: %v", ErrStorageOperation, err)
	}
	return nil
}

func (r *menuRepository) Create(item *models.MenuItem) error {
	r.log.Info("creating menu item", "id", item.ID)

	items, err := r.loadItems()
	if err != nil {
		return err
	}

	*items = append(*items, *item)

	if err := r.saveItems(items); err != nil {
		r.log.Error("failed to save new menu item", "error", err, "id", item.ID)
		return err
	}

	return nil
}

func (r *menuRepository) GetByID(id string) (*models.MenuItem, error) {
	r.log.Info("retrieving menu item", "id", id)

	items, err := r.loadItems()
	if err != nil {
		return nil, err
	}

	for _, item := range *items {
		if item.ID == id {
			itemCopy := item // Create a copy to avoid data races
			return &itemCopy, nil
		}
	}

	return nil, nil
}

func (r *menuRepository) GetAll() (*[]models.MenuItem, error) {
	r.log.Info("retrieving all menu items")

	items, err := r.loadItems()
	if err != nil {
		r.log.Error("failed to load menu items", "error", err)
		return nil, err
	}

	return items, nil
}

func (r *menuRepository) Update(item *models.MenuItem) error {
	r.log.Info("updating menu item", "id", item.ID)

	items, err := r.loadItems()
	if err != nil {
		return err
	}

	for i, existing := range *items {
		if existing.ID == item.ID {
			(*items)[i] = *item

			if err := r.saveItems(items); err != nil {
				r.log.Error("failed to save updated menu item", "error", err, "id", item.ID)
				return err
			}

			break
		}
	}

	return nil
}

func (r *menuRepository) Delete(id string) error {
	r.log.Info("deleting menu item", "id", id)

	items, err := r.loadItems()
	if err != nil {
		return err
	}

	for i, item := range *items {
		if item.ID == id {
			(*items)[i], (*items)[len(*items)-1] = (*items)[len(*items)-1], (*items)[i]
			*items = (*items)[:len(*items)-1]

			if err := r.saveItems(items); err != nil {
				r.log.Error("failed to save updated menu items", "error", err)
				return err
			}

			break
		}
	}

	return nil
}

func (r *menuRepository) GetRequiredIngredients(id string) (*[]models.MenuItemIngredient, error) {
	r.log.Info("retrieving required ingredients", "menu_item_id", id)

	items, err := r.loadItems()
	if err != nil {
		return nil, err
	}

	for _, item := range *items {
		if item.ID == id {
			ingredients := make([]models.MenuItemIngredient, len(item.Ingredients))
			copy(ingredients, item.Ingredients)
			return &ingredients, nil
		}
	}

	return nil, nil
}
