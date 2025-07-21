package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func GetPointsOfInterestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	pointsOfInterest, err := domain.GetPointsOfInterest(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Points Of Interest fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, pointsOfInterest)
}

func PostPointOfInterestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var poi domain.PointOfInterest
	json.NewDecoder(r.Body).Decode(&poi)

	newPoi, err := domain.CreatePointOfInterest(poi.CampaignID, poi.SublocationID, poi.Name)
	if err != nil {
		http.Error(w, "Failed to create point of interest: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newPoi)
}

func PutPointOfInterestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var poi domain.PointOfInterest
	json.NewDecoder(r.Body).Decode(&poi)

	updatedPoi, err := domain.UpdatePointOfInterest(poi.ID, poi.Name)
	if err != nil {
		http.Error(w, "Failed to update point of interest: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updatedPoi)
}

func DeletePointOfInterestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["pointOfInterestId"]
	pointOfInterestID, _ := strconv.Atoi(idStr)

	err := domain.DeletePointOfInterest(pointOfInterestID)
	if err != nil {
		http.Error(w, "Failed to delete point of interest: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
