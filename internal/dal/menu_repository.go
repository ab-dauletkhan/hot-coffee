package dal

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

func GetMenuItemsJSON(r *http.Request) ([]models.MenuItem, error) {
	file, err := os.ReadFile("data/menu_items.json")
	if err != nil {
		return nil, errors.New("couldn't read 'data/menu_items.json'")
	}

	req := []models.MenuItem{}
	if err := json.Unmarshal(file, &req); err != nil {
		return nil, errors.New("couldn't unmarshall menu items")
	}

	return req, nil
}

// func SaveMenuItemsJSON(r *http.Request, menuItems []models.MenuItem) {
// 	data, err := json.MarshalIndent(menuItems, "  ", "  ")
// 	if err != nil {
// 		SaveJSONLog(r, slog.LevelDebug, logCommonFields(r, 500), "couldn't read 'data/menu_items.json'")

// 	}

// 	if err := os.WriteFile("data/menu_items.json", data, 0666); err != nil {
// 		SaveJSONLog(r, slog.LevelDebug, logCommonFields(r, 500), "couldn't read 'data/menu_items.json'")
// 		return
// 	}
// }
