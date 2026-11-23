package api

import (
	"encoding/json"
	"fmt"
	"go_dm_api/auth"
	"go_dm_api/domain"
	"go_dm_api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type completeCampaign struct {
	Campaign             domain.Campaign
	Months               []domain.Month
	WeekDays             []domain.WeekDay
	CalendarCycles       []domain.CalendarCycle
	CalendarCycleOffsets []domain.CampaignCalendarCycleOffset
	CalendarEvents       []domain.CalendarEvent
}

func GetCampaignHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(idStr)

	campaign, err := domain.GetCampaign(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	months, err := domain.GetMonths(campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign months: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	weekDays, err := domain.GetWeekDays(campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign week days: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	calendarCycles, err := domain.GetCalendarCycles(campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign calendar cycles: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	calendarCycleOffsets, err := domain.GetCalendarCycleOffsets(campaign.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign calendar cycle offsets: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	calendarEvents, err := domain.GetCalendarEvents(campaign.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign calendar events: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	utils.StandardResponse(w, completeCampaign{Campaign: *campaign, Months: months, WeekDays: weekDays, CalendarCycles: calendarCycles, CalendarCycleOffsets: calendarCycleOffsets, CalendarEvents: calendarEvents})
}

func ListCampaignHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.ExtractUserId(r)
	campaigns, err := domain.ListCampaigns(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaigns: %s", err.Error()), http.StatusNotFound)
		return
	}

	utils.StandardResponse(w, campaigns)
}

func PostCampaignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.StandardResponse(w, nil)
		return
	}
	utils.EnableCORS(&w)

	var campaign domain.Campaign
	json.NewDecoder(r.Body).Decode(&campaign)

	newCampaign, err := domain.CreateCampaign(campaign.Name, campaign.CurrentPlayerXP, campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating campaign: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	userId := auth.ExtractUserId(r)
	if userId != 0 {
		domain.CreateUserPermission(userId, newCampaign.ID)
	}

	w.WriteHeader(http.StatusCreated)
	utils.StandardResponse(w, newCampaign)
}

func DeleteCampaignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.StandardResponse(w, nil)
		return
	}
	utils.EnableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["campaignId"]
	campaignID, _ := strconv.Atoi(idStr)

	err := domain.DeleteCampaign(campaignID)
	if err != nil {
		http.Error(w, "Failed to delete campaign: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
