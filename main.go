package main

import (
	"yard-planning/routes"
	"yard-planning/app/middleware"
	"net/http"
	"log"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
)
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

    r.Route("/api", api.ApiRoutes())

	log.Println("listening on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}