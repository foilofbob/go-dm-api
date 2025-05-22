package main

import (
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/handlers"
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
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday", api.PostGameDayHandler).Methods("POST", "OPTIONS")

	// Note routes
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{referenceType}", api.GetNotesHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note", api.PostNoteHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{noteId:[0-9]+}", api.PutNoteHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{noteId:[0-9]+}", api.DeleteNoteHandler).Methods("DELETE", "OPTIONS")

	// Players
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/players", api.GetPlayersHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/player", api.PostPlayerHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/player/{playerId:[0-9]+}", api.PutPlayerHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/player/{playerId:[0-9]+}", api.DeletePlayerHandler).Methods("DELETE", "OPTIONS")

	// Items
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/items", api.GetItemsHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/item", api.PostItemHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/item/{itemId:[0-9]+}", api.PutItemHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/item/{itemId:[0-9]+}", api.DeleteItemHandler).Methods("DELETE", "OPTIONS")

	// Start the HTTP server
	log.Println("Server listening on :" + config.Cfg().Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Cfg().Server.Port, r))
}
