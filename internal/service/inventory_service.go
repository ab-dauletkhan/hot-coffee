package service

import (
	"fmt"

	"github.com/ab-dauletkhan/hot-coffee/internal/dal"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func SaveInventoryItem(items []*models.InventoryItem) error {
	inventoryItems := dal.GetJSONInventory()

	for _, item := range items {
		notFound := true

		for i := range inventoryItems {
			if inventoryItems[i].IngredientID == item.IngredientID {
				notFound = false
				if inventoryItems[i].Unit != item.Unit {
					return fmt.Errorf("%s: invalid ingredient unit: have (%s) want (%s)",
						item.Name, item.Unit, inventoryItems[i].Unit)
				}

				inventoryItems[i].Quantity += item.Quantity
			}
		}

		if notFound {
			inventoryItems = append(inventoryItems, item)
		}

	}

	dal.SaveJSONInventoryItem(inventoryItems)
	return nil
}
