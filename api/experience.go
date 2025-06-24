package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func GetExperiencesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	experiences, err := domain.GetExperiences(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Experiences fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, experiences)
}

func PostExperienceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var experience domain.Experience
	json.NewDecoder(r.Body).Decode(&experience)

	newExperience, err := domain.CreateExperience(experience.CampaignID, experience.Description, experience.XP, experience.Finalized)
	if err != nil {
		http.Error(w, "Failed to create experience: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var campaignXP int
	// If finalized, need to update CurrentPlayerXP on the campaign
	if experience.Finalized {
		campaignXP, err = domain.UpdateCurrentPlayerXP(experience.CampaignID, experience.XP)

		if err != nil {
			http.Error(w, "Failed to update campaign XP however new XP has been created: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		campaign, err := domain.GetCampaign(experience.CampaignID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching campaign: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		campaignXP = campaign.CurrentPlayerXP
	}

	responseMap := make(map[string]interface{})
	responseMap["experience"] = newExperience
	responseMap["campaignXP"] = campaignXP

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, responseMap)
}

func PutExperienceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var experience domain.Experience
	json.NewDecoder(r.Body).Decode(&experience)

	updateExperience, err := domain.UpdateExperience(experience.ID, experience.Description, experience.XP, experience.Finalized)
	if err != nil {
		http.Error(w, "Failed to update experience: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var campaignXP int
	// If finalized, need to update CurrentPlayerXP on the campaign
	if experience.Finalized {
		campaignXP, err = domain.UpdateCurrentPlayerXP(experience.CampaignID, experience.XP)

		if err != nil {
			http.Error(w, "Failed to update campaign XP however individual XP has been updated: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		campaign, err := domain.GetCampaign(experience.CampaignID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching campaign: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		campaignXP = campaign.CurrentPlayerXP
	}

	responseMap := make(map[string]interface{})
	responseMap["experience"] = updateExperience
	responseMap["campaignXP"] = campaignXP

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, responseMap)
}

func DeleteExperienceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	// TODO: If finalized, return error

	vars := mux.Vars(r)
	idStr := vars["experienceId"]
	experienceID, _ := strconv.Atoi(idStr)

	err := domain.DeleteExperience(experienceID)
	if err != nil {
		http.Error(w, "Failed to delete experience: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
