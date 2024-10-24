package internal

import (
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/handler"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Order handling
	mux.HandleFunc("POST /orders", handler.CreateOrder)
	mux.HandleFunc("GET /orders", handler.GetAllOrders)
	mux.HandleFunc("GET /orders/{id}", handler.GetOrder)
	mux.HandleFunc("PUT /orders/{id}", handler.PutOrder)
	mux.HandleFunc("DELETE /orders/{id}", handler.DeleteOrder)
	mux.HandleFunc("POST /orders/{id}/close", handler.CloseOrder)

	// Menu handling
	mux.HandleFunc("POST /menu", handler.PostMenu)
	mux.HandleFunc("GET /menu", handler.GetAllMenu)
	mux.HandleFunc("GET /menu/{id}", handler.GetMenu)
	mux.HandleFunc("PUT /menu/{id}", handler.PutMenu)
	mux.HandleFunc("DELETE /menu/{id}", handler.DeleteMenu)

	// Inventory handling
	mux.HandleFunc("POST /inventory", handler.PostInventory)
	mux.HandleFunc("GET /inventory", handler.GetAllInventory)
	mux.HandleFunc("GET /inventory/{id}", handler.GetInventory)
	mux.HandleFunc("PUT /inventory/{id}", handler.PutInventory)
	mux.HandleFunc("DELETE /inventory/{id}", handler.DeleteInventory)

	// Aggregation handling
	mux.HandleFunc("GET /reports/total-sales", handler.GetReportsTotalSales)
	mux.HandleFunc("GET /reports/popular-items", handler.GetReportsPopularItems)

	return mux
}
