package service

import (
	"fmt"

	"github.com/ab-dauletkhan/hot-coffee/internal/dal"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func SaveInventoryItem(item models.InventoryItem) error {
	inventoryItems := dal.GetJSONInventory()

	for i := range inventoryItems {
		if inventoryItems[i].IngredientID == item.IngredientID {
			if inventoryItems[i].Unit != item.Unit {
				return fmt.Errorf("invalid ingredient unit: have (%s) want (%s)",
					item.Unit, inventoryItems[i].Unit)
			}

			inventoryItems[i].Quantity += item.Quantity
			dal.SaveJSONInventoryItem(inventoryItems)
			return nil
		}
	}

	inventoryItems = append(inventoryItems, item)
	dal.SaveJSONInventoryItem(inventoryItems)
	return nil
}
