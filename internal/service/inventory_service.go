package service

import (
	"fmt"

	"github.com/ab-dauletkhan/hot-coffee/internal/dal"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func SaveInventoryItem(i models.InventoryItem) error {
	inventoryItems := dal.GetJSONInventory()

	for _, item := range inventoryItems {
		if item.IngredientID == i.IngredientID {
			if item.Unit != i.Unit {
				return fmt.Errorf("invalid ingredient unit: have (%s) want (%s)",
					i.Unit, item.Unit)
			}

			item.Quantity += i.Quantity
			dal.SaveJSONInventoryItem(inventoryItems)
			return nil
		}
	}

	// TODO: id generation
	inventoryItems = append(inventoryItems, i)
	dal.SaveJSONInventoryItem(inventoryItems)
	return nil
}
