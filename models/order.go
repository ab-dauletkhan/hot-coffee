package models

import (
	"errors"
	"regexp"
	"strings"
)

type Order struct {
	ID           string      `json:"order_id"`
	CustomerName string      `json:"customer_name"`
	Items        []OrderItem `json:"items"`
	Status       string      `json:"status"`
	CreatedAt    string      `json:"created_at"`
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

var validStatus = map[string]bool{
	"pending": true, "completed": true, "canceled": true,
}

var validTimestampRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`)

func (o *Order) IsValid() error {
	if err := o.validateFields(); err != nil {
		return err
	}
	o.normalizeFields()
	return nil
}

func (o *Order) validateFields() error {
	if o.ID == "" || !validIngredientID.MatchString(o.ID) {
		return errors.New("order_id must be non-empty and alphanumeric with underscores only")
	}
	if o.CustomerName == "" || !validNameRegex.MatchString(o.CustomerName) {
		return errors.New("customer_name must contain only letters and spaces")
	}
	if valid := validStatus[o.Status]; !valid {
		return errors.New("status must be one of 'pending', 'completed', or 'canceled'")
	}
	if !validTimestampRegex.MatchString(o.CreatedAt) {
		return errors.New("created_at must be a valid ISO8601 timestamp")
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
