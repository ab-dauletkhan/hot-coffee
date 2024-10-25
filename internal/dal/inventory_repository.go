package dal

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

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
