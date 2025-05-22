package api

import (
	"encoding/json"
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

func PostPlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var player domain.Player
	json.NewDecoder(r.Body).Decode(&player)

	newPlayer, err := domain.CreatePlayer(player.CampaignID, player.Name, player.Race, player.Class, player.ArmorClass, player.HitPoints, player.PassivePerception, player.Languages, player.Movement)
	if err != nil {
		http.Error(w, "Failed to create player: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newPlayer)
}

func PutPlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var player domain.Player
	json.NewDecoder(r.Body).Decode(&player)

	updateItem, err := domain.UpdatePlayer(player.ID, player.Name, player.Race, player.Class, player.ArmorClass, player.HitPoints, player.PassivePerception, player.Languages, player.Movement)
	if err != nil {
		http.Error(w, "Failed to update player: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updateItem)
}

func DeletePlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["playerId"]
	playerID, _ := strconv.Atoi(idStr)

	err := domain.DeletePlayer(playerID)
	if err != nil {
		http.Error(w, "Failed to delete player: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
