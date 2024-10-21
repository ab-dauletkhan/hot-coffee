package main

import (
	"log"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/handler"
)

func main() {
	http.HandleFunc("/orders/", handler.OrdersHandler)
	http.HandleFunc("/menu/", handler.MenuHandler)
	http.HandleFunc("/inventory/", handler.InventoryHandler)
	http.HandleFunc("/reports/", handler.ReportsHandler)

	port := ":8080"
	log.Printf("Starting development server at http://localhost%s/\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
