package dal

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

func GetJSONInventory() []models.InventoryItem {
	data, err := os.ReadFile("data/inventory.json")
	if err != nil {
		slog.Debug(fmt.Sprintf("error reading inventory.json: %v", err))
	}

	inventoryItems := []models.InventoryItem{}

	if err := json.Unmarshal(data, &inventoryItems); err != nil {
		slog.Debug(fmt.Sprintf("error unmarshalling inventory items: %v", err))
	}

	return inventoryItems
}

func SaveJSONInventoryItem(items []models.InventoryItem) {
	data, err := json.MarshalIndent(items, "  ", "  ")
	if err != nil {
		slog.Debug(fmt.Sprintf("error marshalling inventory items: %v", err))
	}

	if err := os.WriteFile("data/inventory.json", data, filePerm); err != nil {
		slog.Debug(fmt.Sprintf("error writing to inventory.json: %v", err))
	}
}
