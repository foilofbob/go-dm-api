package api

import (
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

func GetLocationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	locationIdStr := vars["locationId"]

	locationId, err := strconv.Atoi(locationIdStr)

	location, err := domain.GetLocation(locationId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Location fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	
	StandardResponse(w, location)
}

// Post

// Update

// Delete
