package api

import (
	"encoding/json"
	"net/http"
)

func StandardResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow", "Accept, Cache-Control, Content-Type, Origin, User-Agent")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Cache-Control, Content-Type, Origin, User-Agent")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	json.NewEncoder(w).Encode(v)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Cache-Control, Content-Type, Origin, User-Agent")
}
