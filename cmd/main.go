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

	log.Println("Starting development server at http://localhosts:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
