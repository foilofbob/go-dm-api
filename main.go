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
	r.HandleFunc("/campaign", api.PostCampaignHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{id:[0-9]+}", api.GetCampaignHandler).Methods("GET")

	// CampaignSetting routes
	r.HandleFunc("/campaign-settings", api.ListCampaignSettingHandler).Methods("GET")

	// Categories
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/categories", api.GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/category", api.PostCategoryHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/category/{categoryId:[0-9]+}", api.PutCategoryHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/category/{categoryId:[0-9]+}", api.DeleteCategoryHandler).Methods("DELETE", "OPTIONS")

	// Characters
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/players", api.GetPlayersHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/npcs", api.GetNPCsHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/character", api.PostCharacterHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/character/{characterId:[0-9]+}", api.PutCharacterHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/character/{characterId:[0-9]+}", api.DeleteCharacterHandler).Methods("DELETE", "OPTIONS")

	// Experience
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/experiences", api.GetExperiencesHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/experience", api.PostExperienceHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/experience/{experienceId:[0-9]+}", api.PutExperienceHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/experience/{experienceId:[0-9]+}", api.DeleteExperienceHandler).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/experience/clear-finalized", api.DeleteExperienceClearFinalizedHandler).Methods("DELETE", "OPTIONS")

	// GameDay routes
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday", api.GetGameDaysHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday", api.PostGameDayHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday/initialize", api.PostInitializeGameDayHandler).Methods("POST", "OPTIONS")

	// Items
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/items", api.GetItemsHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/item", api.PostItemHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/item/{itemId:[0-9]+}", api.PutItemHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/item/{itemId:[0-9]+}", api.DeleteItemHandler).Methods("DELETE", "OPTIONS")

	// Locations
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/locations", api.GetLocationsHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/location", api.PostLocationHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/location/{locationId:[0-9]+}", api.PutLocationHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/location/{locationId:[0-9]+}", api.DeleteLocationHandler).Methods("DELETE", "OPTIONS")

	// Note routes
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{referenceType}", api.GetNotesHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note", api.PostNoteHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{noteId:[0-9]+}", api.PutNoteHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{noteId:[0-9]+}", api.DeleteNoteHandler).Methods("DELETE", "OPTIONS")

	// Points of interest
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/points-of-interest", api.GetPointsOfInterestHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/point-of-interest", api.PostPointOfInterestHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/point-of-interest/{pointOfInterestId:[0-9]+}", api.PutPointOfInterestHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/point-of-interest/{pointOfInterestId:[0-9]+}", api.DeletePointOfInterestHandler).Methods("DELETE", "OPTIONS")

	// Spell Books
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbooks", api.GetSpellBooksHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook", api.PostSpellBookHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}", api.GetSpellBookHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}", api.DeleteSpellBookHandler).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}", api.PutSpellBookHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}/spell", api.PostSpellBookEntryHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}/spell/{spellBookEntryId}", api.DeleteSpellBookEntryHandler).Methods("DELETE", "OPTIONS")

	// Spells
	r.HandleFunc("/spells", api.ListSpellsHandler).Methods("GET")

	// Sublocations
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/sublocations", api.GetSublocationsHandler).Methods("GET")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/sublocation", api.PostSublocationHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/sublocation/{sublocationId:[0-9]+}", api.PutSublocationHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/campaign/{campaignId:[0-9]+}/sublocation/{sublocationId:[0-9]+}", api.DeleteSublocationHandler).Methods("DELETE", "OPTIONS")

	// Start the HTTP server
	log.Println("Server listening on :" + config.Cfg().Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Cfg().Server.Port, r))
}
