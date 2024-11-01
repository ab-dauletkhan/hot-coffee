package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	validNameRegex    = regexp.MustCompile(`^[A-Za-z ]+$`)
	validIngredientID = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
)

type InventoryItem struct {
	IngredientID string  `json:"ingredient_id"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
}

// IsValid performs validation and normalization on the InventoryItem.
func (i *InventoryItem) IsValid() error {
	if err := i.validateFields(); err != nil {
		return err
	}
	i.normalizeFields()
	return nil
}

func (i *InventoryItem) validateFields() error {
	if i.IngredientID == "" || !validIngredientID.MatchString(i.IngredientID) {
		return errors.New("ingredient_id must be non-empty and alphanumeric with underscores only")
	}
	if i.Name == "" {
		return errors.New("inventory name cannot be empty")
	}
	if !validNameRegex.MatchString(i.Name) {
		return errors.New("name must contain only letters and spaces")
	}
	if i.Quantity <= 0 {
		return errors.New("quantity should be bigger than 0")
	}
	if !isValidUnit(i.Unit) {
		return fmt.Errorf("unit must be one of %v", validUnitsString())
	}
	return nil
}

func (i *InventoryItem) normalizeFields() {
	i.Name = strings.Title(strings.TrimSpace(i.Name))
	i.Unit = strings.ToLower(strings.TrimSpace(i.Unit))

	// i.IngredientID = generateID(i.Name)
}

func generateID(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "_"))
}

var validUnits = map[string]bool{
	"g":     true,
	"ml":    true,
	"shots": true,
}

func validUnitsString() string {
	var units []string
	for unit := range validUnits {
		units = append(units, unit)
	}
	return strings.Join(units, ", ")
}

func isValidUnit(unit string) bool {
	return validUnits[strings.ToLower(unit)]
}
