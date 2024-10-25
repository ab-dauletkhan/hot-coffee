package models

import (
	"errors"
	"regexp"
	"strings"
)

var (
	validNameRegex = regexp.MustCompile(`^[A-Za-z ]+$`)
)

type InventoryItem struct {
	IngredientID string  `json:"ingredient_id"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
}

func (i *InventoryItem) IsValid() error {
	if err := i.validateFields(); err != nil {
		return err
	}
	i.normalizeFields()
	return nil
}

func (i *InventoryItem) validateFields() error {
	if i.Name == "" || i.Unit == "" || i.Quantity == 0 {
		return errors.New("all fields (name, unit, quantity) must be provided")
	}
	if i.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if !validNameRegex.MatchString(i.Name) {
		return errors.New("name must contain only letters and spaces")
	}
	if !isValidUnit(i.Unit) {
		return errors.New("unit can only be 'ml', 'g', or 'shots'")
	}
	return nil
}

func (i *InventoryItem) normalizeFields() {
	i.Name = strings.Title(strings.TrimSpace(i.Name))
	i.IngredientID = generateIDFromName(i.Name)
}

func isValidUnit(unit string) bool {
	validUnits := map[string]struct{}{
		"g": {}, "ml": {}, "shots": {},
	}
	_, valid := validUnits[unit]
	return valid
}

func generateIDFromName(name string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(name)), " ", "_")
}
