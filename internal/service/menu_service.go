package service

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/dal"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func GetMenuItemsJSON(r *http.Request) ([]models.MenuItem, error) {
	menuItems, err := dal.GetMenuItemsJSON(r)
	if err != nil {
		CreateLog(r, slog.LevelDebug, 500, fmt.Sprint(err))
		return nil, err
	}

	return menuItems, nil
}

func SaveMenuItemsJSON(r *http.Request, menuItems []models.MenuItem) error {
	return nil
}
