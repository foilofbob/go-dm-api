package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

/*
TODO: Eventually this will need to be paginated, and the response will need to include the relevant notes
- fetch in batches based on page number
- fetch notes based on those ids
- new struct to contain days, notes, and pagination details
*/
func GetGameDaysHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	gameDays, err := domain.GetGameDays(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Game Days fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, gameDays)
}

// PostInitializeGameDayHandler This endpoint is for creating the first game day of a campaign
func PostInitializeGameDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var initializeGameDay domain.InitializeGameDay
	json.NewDecoder(r.Body).Decode(&initializeGameDay)
	gameDay := initializeGameDay.GameDay

	newGameDay, err := domain.CreateGameDay(gameDay.CampaignID, gameDay.InGameDay, gameDay.Day, gameDay.Month, gameDay.Year)
	if err != nil {
		http.Error(w, "Failed to create initial Game Day", http.StatusInternalServerError)
		return
	}

	campaign, err := domain.GetCampaign(gameDay.CampaignID)
	if err != nil {
		http.Error(w, fmt.Sprintf("DB error reading campaign: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	calendarCycles, err := domain.GetCalendarCycles(campaign.CampaignSettingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("DB error reading campaign calendar cycles: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	for _, cycle := range calendarCycles {
		offset := 0

		for _, submittedCycle := range initializeGameDay.Cycles {
			if submittedCycle.ID == cycle.ID {
				offset = submittedCycle.Offset
			}
		}

		domain.CreateCalendarCycleOffset(gameDay.CampaignID, cycle.ID, offset)
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newGameDay)
}

// PostGameDayHandler For a campaign that is already up and running, this endpoint adds the next consecutive day
func PostGameDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["campaignId"]
	campaignID, _ := strconv.Atoi(idStr)

	mostRecentGameDay, err := domain.GetMostRecentGameDay(campaignID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create Game Day: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	campaignMonthSummary, err := domain.GetCampaignMonthSummary(campaignID, mostRecentGameDay.Month)

	day := mostRecentGameDay.Day + 1
	month := mostRecentGameDay.Month
	year := mostRecentGameDay.Year

	if day > campaignMonthSummary.CurrentMonthDays {
		month += 1
		day = 1

		if month > campaignMonthSummary.MonthCount {
			month = 1
			year += 1
		}
	}

	newGameDay, err := domain.CreateGameDay(mostRecentGameDay.CampaignID, mostRecentGameDay.InGameDay+1, day, month, year)
	if err != nil {
		http.Error(w, "Failed to create Game Day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newGameDay)
}
