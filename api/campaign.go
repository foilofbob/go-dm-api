package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/config"
	"go_dm_api/domain"
	"log"
	"net/http"
	"strconv"
)

func GetCampaignHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(config.DBDriver(), config.DBConnectString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	campaignId, err := strconv.Atoi(idStr)

	// Call the GetUser function to fetch the user data from the database
	campaign, err := domain.GetCampaign(db, campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Campaign not found: %s", err.Error()), http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(campaign)
}

func ListCampaignHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("Fetching campaigns"))
	db, err := sql.Open(config.DBDriver(), config.DBConnectString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Call the GetUser function to fetch the user data from the database
	campaigns, err := domain.ListCampaigns(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaigns: %s", err.Error()), http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8280")
	json.NewEncoder(w).Encode(campaigns)
}
