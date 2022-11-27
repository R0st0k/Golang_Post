package main

import (
	post "backend/api"
	"backend/db"
	"github.com/rs/cors"
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

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", srv))

	handler := cors.Handler(mux)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
