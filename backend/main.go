package main

import (
	post "backend/api"
	"backend/db"
	"log"
	"net/http"
)

func main() {
	db.Init()

	// Create service instance.
	service := &postService{}

	// Create generated server.
	srv, err := post.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", srv))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
