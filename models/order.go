package models

import (
	"errors"
	"strings"
)

type Order struct {
	ID           string      `json:"order_id,omitempty"`
	CustomerName string      `json:"customer_name"`
	Items        []OrderItem `json:"items"`
	Status       string      `json:"status,omitempty"`
	CreatedAt    string      `json:"created_at,omitempty"`
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

var (
	ErrItemNotAvailable = errors.New("ingridient not available")

	StatusPending   = "pending"
	StatusCompleted = "closed"
)

func (o *Order) IsValid() error {
	if err := o.validateFields(); err != nil {
		return err
	}
	o.normalizeFields()
	return nil
}

func (o *Order) validateFields() error {
	if o.ID != "" {
		return errors.New("order_id must not be provided")
	}
	if o.CustomerName == "" || !validNameRegex.MatchString(o.CustomerName) {
		return errors.New("customer_name must contain only letters and spaces")
	}
	for _, item := range o.Items {
		if err := item.IsValid(); err != nil {
			return err
		}
	}
	return nil
}

func (o *Order) normalizeFields() {
	o.CustomerName = strings.Title(strings.TrimSpace(o.CustomerName))
}

func (oi *OrderItem) IsValid() error {
	if oi.ProductID == "" || !validIngredientID.MatchString(oi.ProductID) {
		return errors.New("product_id must be non-empty and alphanumeric with underscores only")
	}
	if oi.Quantity <= 0 {
		return errors.New("quantity must be a positive integer")
	}
	return nil
}
