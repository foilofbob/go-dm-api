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
	Campaign             domain.Campaign
	Months               []domain.Month
	WeekDays             []domain.WeekDay
	CalendarCycles       []domain.CalendarCycle
	CalendarCycleOffsets []domain.CampaignCalendarCycleOffset
	CalendarEvents       []domain.CalendarEvent
}

func GetCampaignHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

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

	StandardResponse(w, completeCampaign{Campaign: *campaign, Months: months, WeekDays: weekDays, CalendarCycles: calendarCycles, CalendarCycleOffsets: calendarCycleOffsets, CalendarEvents: calendarEvents})
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
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var campaign domain.Campaign
	json.NewDecoder(r.Body).Decode(&campaign)

	newCampaign, err := domain.CreateCampaign(campaign.Name, campaign.CurrentPlayerXP, campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating campaign: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newCampaign)
}
