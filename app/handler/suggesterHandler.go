package handler

import (
	"encoding/json"
	"net/http"
)

func CreateSuggest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	code := http.StatusOK
	v := map[string]string{"message": "Hello, Suggestion"}
	
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}