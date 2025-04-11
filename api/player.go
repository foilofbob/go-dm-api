package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func GetPlayersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	players, err := domain.GetPlayers(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Players fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, players)
}
