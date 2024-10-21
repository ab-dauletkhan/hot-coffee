package main

import (
	"log"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/router"
)

func main() {
	// TODO: Implement "--help"

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Routes(),
	}

	log.Println("Starting development server at http://localhosts:8080")
	log.Fatal(srv.ListenAndServe())
}
