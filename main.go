package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go_dm_api/api"
	"go_dm_api/config"
	"log"
	"net/http"
)

func main() {
	// Initialize config
	config.InitConfig()

	// Create a new router
	r := mux.NewRouter()

	// Campaign routes
	r.HandleFunc("/campaign", api.ListCampaignHandler).Methods("GET")
	r.HandleFunc("/campaign", api.PostCampaignHandler).Methods("POST")
	r.HandleFunc("/campaign/{id:[0-9]+}", api.GetCampaignHandler).Methods("GET")

	// GameDay routes
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday", api.GetGameDaysHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday", api.PostGameDayHandler).Methods("POST")

	// Note routes
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{referenceType}", api.GetNotesHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note", api.PostNoteHandler).Methods("POST")

	// Start the HTTP server
	log.Println("Server listening on :" + config.Cfg().Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Cfg().Server.Port, r))
}
