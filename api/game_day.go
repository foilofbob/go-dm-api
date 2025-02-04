package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

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

	var gameDay domain.GameDay
	json.NewDecoder(r.Body).Decode(&gameDay)

	newGameDay, err := domain.CreateGameDay(mostRecentGameDay.CampaignID, mostRecentGameDay.InGameDay+1, day, month, year)
	if err != nil {
		http.Error(w, "Failed to create Game Day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newGameDay)
}
