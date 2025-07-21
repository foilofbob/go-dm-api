package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func GetSublocationsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	sublocations, err := domain.GetSublocations(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Sublocations fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, sublocations)
}

func PostSublocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var sublocation domain.Sublocation
	json.NewDecoder(r.Body).Decode(&sublocation)

	newSublocation, err := domain.CreateSublocation(sublocation.CampaignID, sublocation.LocationID, sublocation.Name, sublocation.Description)
	if err != nil {
		http.Error(w, "Failed to create sublocation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newSublocation)
}

func PutSublocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var sublocation domain.Sublocation
	json.NewDecoder(r.Body).Decode(&sublocation)

	updatedSublocation, err := domain.UpdateSublocation(sublocation.ID, sublocation.Name, sublocation.Description)
	if err != nil {
		http.Error(w, "Failed to update sublocation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updatedSublocation)
}

func DeleteSublocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["sublocationId"]
	sublocationID, _ := strconv.Atoi(idStr)

	err := domain.DeleteSublocation(sublocationID)
	if err != nil {
		http.Error(w, "Failed to delete sublocation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
