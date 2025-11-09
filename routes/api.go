package api

import (
	"net/http"
	"encoding/json"
	
	"github.com/go-chi/chi/v5"
)

func ApiRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "World"
			}
			writeJSON(w, http.StatusOK, map[string]string{"message": "Hello, " + name})
		})

		r.Get("/suggestion", func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, http.StatusOK, map[string]string{"message": "Hello, Suggestion"})
		})

		r.Get("/placement", func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, http.StatusOK, map[string]string{"message": "Hello, Suggestion"})
		})

		r.Get("/pickup", func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, http.StatusOK, map[string]string{"message": "Hello, Suggestion"})
		})
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}