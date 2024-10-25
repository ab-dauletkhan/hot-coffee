package service

import (
	"fmt"

	"github.com/ab-dauletkhan/hot-coffee/internal/dal"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

func SaveInventoryItem(items []*models.InventoryItem) error {
	inventoryItems := dal.GetJSONInventory()
	itemMap := mapInventoryItems(inventoryItems)

	for _, item := range items {
		if err := updateOrAddItem(item, itemMap); err != nil {
			return err
		}
	}

	dal.SaveJSONInventoryItem(mapToSlice(itemMap))
	return nil
}

func mapInventoryItems(items []*models.InventoryItem) map[string]*models.InventoryItem {
	itemMap := make(map[string]*models.InventoryItem)
	for _, item := range items {
		itemMap[item.IngredientID] = item
	}
	return itemMap
}

func updateOrAddItem(item *models.InventoryItem, itemMap map[string]*models.InventoryItem) error {
	if existingItem, exists := itemMap[item.IngredientID]; exists {
		if existingItem.Unit != item.Unit {
			return fmt.Errorf("%s: invalid ingredient unit - have (%s), want (%s)",
				item.Name, item.Unit, existingItem.Unit)
		}
		existingItem.Quantity += item.Quantity
	} else {
		itemMap[item.IngredientID] = item
	}
	return nil
}

func mapToSlice(itemMap map[string]*models.InventoryItem) []*models.InventoryItem {
	items := make([]*models.InventoryItem, 0, len(itemMap))
	for _, item := range itemMap {
		items = append(items, item)
	}
	return items
}
