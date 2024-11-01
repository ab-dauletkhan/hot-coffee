package repository

import (
	"log/slog"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

type MenuRepository interface {
	Create(item *models.MenuItem) error
	GetByID(id string) (*models.MenuItem, error)
	GetAll() ([]*models.MenuItem, error)
	Update(item *models.MenuItem) error
	Delete(id string) error
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

func (r *menuRepository) Create(item *models.MenuItem) error {
	return nil
}

func (r *menuRepository) GetByID(id string) (*models.MenuItem, error) {
	return nil, nil
}

func (r *menuRepository) GetAll() ([]*models.MenuItem, error) {
	return nil, nil
}

func (r *menuRepository) Update(item *models.MenuItem) error {
	return nil
}

func (r *menuRepository) Delete(id string) error {
	return nil
}

// const menuItemsPath = "menu_items.json"

// func GetJSONMenuItems(r *http.Request) ([]models.MenuItem, error) {
// 	file, err := os.ReadFile(filepath.Join(core.Dir, menuItemsPath))
// 	if err != nil {
// 		slog.Debug(fmt.Sprintf("error reading menu_items.json: %v", err))
// 		return []models.MenuItem{}, fmt.Errorf("failed to read menu items: %w", err)
// 	}

// 	req := []models.MenuItem{}
// 	if err := json.Unmarshal(file, &req); err != nil {
// 		slog.Debug(fmt.Sprintf("error unmarshalling menu items: %v", err))
// 		return []models.MenuItem{}, fmt.Errorf("failed to parse menu items: %w", err)
// 	}

// 	return req, nil
// }

// func SaveJSONMenuItem(items []models.MenuItem) error {
// 	data, err := json.MarshalIndent(items, "", "  ")
// 	if err != nil {
// 		slog.Debug(fmt.Sprintf("error marshalling menu items: %v", err))
// 		return fmt.Errorf("failed to save menu items: %w", err)
// 	}

// 	if err := os.WriteFile(filepath.Join(core.Dir, menuItemsPath), data, core.DirPerm); err != nil {
// 		slog.Debug(fmt.Sprintf("error writing to menu_items.json: %v", err))
// 		return fmt.Errorf("failed to write menu items: %w", err)
// 	}

// 	return nil
// }

// // 	if err := os.WriteFile("data/menu_items.json", data, 0666); err != nil {
// // 		SaveJSONLog(r, slog.LevelDebug, logCommonFields(r, 500), "couldn't read 'data/menu_items.json'")
// // 		return
// // 	}
// // }
