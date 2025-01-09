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
	var gameDay domain.GameDay
	json.NewDecoder(r.Body).Decode(&gameDay)

	err := domain.CreateGameDay(gameDay.CampaignID, gameDay.InGameDay, gameDay.Day, gameDay.Month, gameDay.Year)
	if err != nil {
		http.Error(w, "Failed to create Game Day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Game Day created successfully")
}
