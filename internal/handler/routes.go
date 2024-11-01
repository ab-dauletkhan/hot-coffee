package handler

import "net/http"

func Routes(orderHandler *OrderHandler, menuHandler *MenuHandler, inventoryHandler *InventoryHandler) *http.ServeMux {
	// Setup router (using standard net/http for example)
	mux := http.NewServeMux()

	// ========================
	// Order routes
	// ========================
	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			orderHandler.CreateOrder(w, r)
		case http.MethodGet:
			orderHandler.GetAllOrders(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orderHandler.GetOrder(w, r)
		case http.MethodPut:
			orderHandler.PutOrder(w, r)
		case http.MethodDelete:
			orderHandler.DeleteOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/orders/{id}/close", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			orderHandler.CloseOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// ========================
	// Menu routes
	// ========================
	mux.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			menuHandler.CreateMenuItem(w, r)
		case http.MethodGet:
			menuHandler.GetAllMenu(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/menu/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			menuHandler.GetMenuItem(w, r)
		case http.MethodPut:
			menuHandler.PutMenuItem(w, r)
		case http.MethodDelete:
			menuHandler.DeleteMenuItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// ========================
	// Inventory routes
	// ========================
	mux.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			inventoryHandler.AddInventoryItem(w, r)
		case http.MethodGet:
			inventoryHandler.GetAllInventory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/inventory/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			inventoryHandler.GetInventory(w, r)
		case http.MethodPut:
			inventoryHandler.PutInventory(w, r)
		case http.MethodDelete:
			inventoryHandler.DeleteInventory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}
