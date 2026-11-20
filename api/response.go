package api

import (
	"encoding/json"
	"fmt"
	"go_dm_api/config"
	"go_dm_api/domain"
	"net/http"
)

type DataWithNotes struct {
	Data  any
	Notes []domain.Note
}

func StandardResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow", "Accept, Cache-Control, Content-Type, Origin, User-Agent")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Cache-Control, Content-Type, Origin, User-Agent")
	w.Header().Set("Access-Control-Allow-Origin", config.Origin())
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Vary", "Origin")
	json.NewEncoder(w).Encode(v)
}

func StandardResponseWithNotes(w http.ResponseWriter, v any, campaignId int, referenceType string) {
	notes, err := domain.GetNotes(campaignId, referenceType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Notes read failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	resp := DataWithNotes{Data: v, Notes: notes}

	StandardResponse(w, resp)
}

func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", config.Origin())
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Cache-Control, Content-Type, Origin, User-Agent")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Vary", "Origin")
}
