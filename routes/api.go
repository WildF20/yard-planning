package api

import (
	"yard-planning/app/handler"
	
	"github.com/go-chi/chi/v5"
)

func ApiRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/suggestion", handler.CreateSuggest)
		r.Get("/placement", handler.CreatePlacement)
		r.Get("/pickup", handler.DispatchSlot)
	}
}