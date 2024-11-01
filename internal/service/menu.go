package service

import (
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/internal/repository"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

type MenuService interface {
	CreateMenuItem(item *models.MenuItem) error
	GetMenuItem(id string) (*models.MenuItem, error)
	GetAllMenuItems() ([]*models.MenuItem, error)
	UpdateMenuItem(item *models.MenuItem) error
	DeleteMenuItem(id string) error
	GetPopularItems() ([]*models.MenuItem, error)
}

// MenuService handles business logic for menu items
type menuService struct {
	menuRepo repository.MenuRepository
	log      *slog.Logger
}

// NewMenuService initializes MenuService with repository and logging
func NewMenuService(menuRepo repository.MenuRepository, log *slog.Logger) menuService {
	return menuService{
		menuRepo: menuRepo,
		log:      log,
	}
}

func (s menuService) CreateMenuItem(item *models.MenuItem) error {
	return nil
}

func (s menuService) GetMenuItem(id string) (*models.MenuItem, error) {
	return nil, nil
}

func (s menuService) GetAllMenuItems() ([]*models.MenuItem, error) {
	return nil, nil
}

func (s menuService) UpdateMenuItem(item *models.MenuItem) error {
	return nil
}

func (s menuService) DeleteMenuItem(id string) error {
	return nil
}

func (s menuService) GetPopularItems() ([]*models.MenuItem, error) {
	return nil, nil
}
