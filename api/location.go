package api

import (
	"encoding/json"
	"fmt"
	"go_dm_api/domain"
	"go_dm_api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	utils.StandardResponseWithNotes(w, locations, campaignId, "LOCATION")
}

func PostLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.StandardResponse(w, nil)
		return
	}
	utils.EnableCORS(&w)

	var location domain.Location
	json.NewDecoder(r.Body).Decode(&location)

	newLocation, err := domain.CreateLocation(location.CampaignID, location.Name)
	if err != nil {
		http.Error(w, "Failed to create location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.StandardResponse(w, newLocation)
}

func PutLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.StandardResponse(w, nil)
		return
	}
	utils.EnableCORS(&w)

	var location domain.Location
	json.NewDecoder(r.Body).Decode(&location)

	updatedLocation, err := domain.UpdateLocation(location.ID, location.Name)
	if err != nil {
		http.Error(w, "Failed to update location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.StandardResponse(w, updatedLocation)
}

func DeleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.StandardResponse(w, nil)
		return
	}
	utils.EnableCORS(&w)

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
