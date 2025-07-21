package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	locations, err := domain.GetLocations(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Locations fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, locations)
}

func PostLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var location domain.Location
	json.NewDecoder(r.Body).Decode(&location)

	newLocation, err := domain.CreateLocation(location.CampaignID, location.Name)
	if err != nil {
		http.Error(w, "Failed to create location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newLocation)
}

func PutLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var location domain.Location
	json.NewDecoder(r.Body).Decode(&location)

	updatedLocation, err := domain.UpdateLocation(location.ID, location.Name)
	if err != nil {
		http.Error(w, "Failed to update location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updatedLocation)
}

func DeleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["locationId"]
	locationID, _ := strconv.Atoi(idStr)

	err := domain.DeleteLocation(locationID)
	if err != nil {
		http.Error(w, "Failed to delete location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
