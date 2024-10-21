package main

import (
	"log"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/router"
)

func main() {
	// TODO: Implement flag parsing (e.g. --help, --port)
	// Exit on invalid flags (e.g., invalid command-line arguments, failure to bind to a port)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Routes(),
	}

	log.Println("Starting development server at http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
