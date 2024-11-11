package models

import (
	"errors"
	"fmt"
	"strings"
)

type MenuItem struct {
	ID          string               `json:"product_id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Price       float64              `json:"price"`
	Ingredients []MenuItemIngredient `json:"ingredients"`
}

type MenuItemIngredient struct {
	IngredientID string  `json:"ingredient_id"`
	Quantity     float64 `json:"quantity"`
}

func (m *MenuItem) IsValid() error {
	if err := m.validateFields(); err != nil {
		return err
	}
	m.normalizeFields()
	return nil
}

func (m *MenuItem) validateFields() error {
	if m.ID == "" || !validIngredientID.MatchString(m.ID) {
		return errors.New("product_id must be non-empty and alphanumeric with underscores only")
	}
	if m.Name == "" || !validNameRegex.MatchString(m.Name) {
		return errors.New("name must contain only letters and spaces")
	}
	if m.Price < 0 {
		return errors.New("price must be non-negative")
	}
	if len(m.Description) > 500 {
		return errors.New("description cannot exceed 500 characters")
	}
	for _, ingredient := range m.Ingredients {
		if err := ingredient.IsValid(); err != nil {
			return fmt.Errorf("invalid ingredient in menu item: %w", err)
		}
	}
	return nil
}

func (m *MenuItem) normalizeFields() {
	m.Name = strings.Title(strings.TrimSpace(m.Name))
	m.Description = strings.TrimSpace(m.Description)
}

func (mi *MenuItemIngredient) IsValid() error {
	if mi.IngredientID == "" {
		return errors.New("ingredient_id must be non-empty and alphanumeric with underscores only")
	}
	if mi.Quantity <= 0 {
		return errors.New("quantity must be a positive number")
	}
	return nil
}
