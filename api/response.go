package api

import (
	"encoding/json"
	"net/http"
)

func StandardResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8280")
	json.NewEncoder(w).Encode(v)
}
