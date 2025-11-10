package handler

import (
	"encoding/json"
	"net/http"
)

func CreatePlacement(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	code := http.StatusOK
	v := map[string]string{"message": "Hello, Placement"}
	
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func DispatchSlot(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	code := http.StatusOK
	v := map[string]string{"message": "Hello, Dispatch"}
	
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}