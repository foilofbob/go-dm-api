package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

type completeCampaign struct {
	Campaign domain.Campaign
	Months   []domain.Month
	WeekDays []domain.WeekDay
}

func GetCampaignHandler(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	campaignId, err := strconv.Atoi(idStr)

	campaign, err := domain.GetCampaign(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign: %s", err.Error()), http.StatusNotFound)
		return
	}

	months, err := domain.GetMonths(campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign months: %s", err.Error()), http.StatusNotFound)
		return
	}

	weekDays, err := domain.GetWeekDays(campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign week days: %s", err.Error()), http.StatusNotFound)
		return
	}

	StandardResponse(w, completeCampaign{Campaign: *campaign, Months: months, WeekDays: weekDays})
}

func ListCampaignHandler(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := domain.ListCampaigns()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaigns: %s", err.Error()), http.StatusNotFound)
		return
	}

	StandardResponse(w, campaigns)
}

func PostCampaignHandler(w http.ResponseWriter, r *http.Request) {
	var campaign domain.Campaign
	json.NewDecoder(r.Body).Decode(&campaign)

	err := domain.CreateCampaign(campaign.Name, campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaigns: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Campaign created successfully")
}
