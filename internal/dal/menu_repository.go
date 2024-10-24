package dal

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

func GetJSONMenuItems(r *http.Request) []models.MenuItem {
	file, err := os.ReadFile("data/menu_items.json")
	if err != nil {
		slog.Debug(fmt.Sprintf("error reading menu_items.json: %v", err))
	}

	req := []models.MenuItem{}
	if err := json.Unmarshal(file, &req); err != nil {
		slog.Debug(fmt.Sprintf("error unmarshalling menu items: %v", err))
	}

	return req
}

// func SaveJSONMenuItems(r *http.Request, menuItems []models.MenuItem) {
// 	data, err := json.MarshalIndent(menuItems, "  ", "  ")
// 	if err != nil {
// 		SaveJSONLog(r, slog.LevelDebug, logCommonFields(r, 500), "couldn't read 'data/menu_items.json'")

// 	}

// 	if err := os.WriteFile("data/menu_items.json", data, 0666); err != nil {
// 		SaveJSONLog(r, slog.LevelDebug, logCommonFields(r, 500), "couldn't read 'data/menu_items.json'")
// 		return
// 	}
// }
