package main

import (
	"log"
	"net/http"
)

func main() {
	// TODO: Implement "--help"

	log.Println("Starting development server at http://localhosts:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
