package v1

import (
	"encoding/json"
	"net/http"
)

func ReturnJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func ReturnError(w http.ResponseWriter, code int, message string) {
	ReturnJSON(w, code, map[string]string{"error": message})
}
