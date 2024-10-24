package models

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	ValidNameRegex = regexp.MustCompile(`^[A-Za-z ]+$`)
	ValidIDRegex   = regexp.MustCompile("^([a-z0-9_]{3,63})$")
)

type InventoryItem struct {
	IngredientID string  `json:"ingredient_id"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
}

func (i *InventoryItem) IsValid() error {
	if len(i.Name) == 0 || len(i.Unit) == 0 || i.Quantity == 0 {
		return fmt.Errorf("all fields must be filled")
	}

	if i.Quantity < 0 {
		return fmt.Errorf("quantity cannot be negative")
	}

	i.Name = strings.Title(i.Name)
	i.IngredientID = convertToId(i.Name)
	if !ValidNameRegex.MatchString(i.Name) {
		return fmt.Errorf("name must contain only letters and a space")
	}

	if i.Unit != "g" && i.Unit != "ml" && i.Unit != "shots" {
		return fmt.Errorf("unit can only be in 'ml', 'g', or 'shots'")
	}

	return nil
}

func convertToId(s string) string {
	result := []byte(strings.ToLower(s))

	for i := range result {
		if result[i] == ' ' {
			result[i] = '_'
		}
	}

	return string(result)
}
