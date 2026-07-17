package main

import (
	"go_dm_api/api"
	"go_dm_api/auth"
	"go_dm_api/config"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// After adding all the backend auth, frontend just needs to include {:with-credentials? true} in theory

func main() {
	config.InitConfig()
	r := mux.NewRouter()

	r.HandleFunc("/login", auth.HandleLogin).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", auth.HandleLogout).Methods("POST", "OPTIONS")

	authenticated := r.PathPrefix("/").Subrouter()
	authenticated.Use(auth.UserAccessMiddleware)

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(auth.CampaignAccessMiddleware)

	// Campaign routes
	authenticated.HandleFunc("/campaign", api.ListCampaignHandler).Methods("GET")
	authenticated.HandleFunc("/campaign", api.PostCampaignHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}", api.GetCampaignHandler).Methods("GET")

	// CampaignSetting routes
	authenticated.HandleFunc("/campaign-settings", api.ListCampaignSettingHandler).Methods("GET")

	// Categories
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/categories", api.GetCategoriesHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/category", api.PostCategoryHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/category/{categoryId:[0-9]+}", api.PutCategoryHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/category/{categoryId:[0-9]+}", api.DeleteCategoryHandler).Methods("DELETE", "OPTIONS")

	// Characters
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/players", api.GetPlayersHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/npcs", api.GetNPCsHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/character", api.PostCharacterHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/character/{characterId:[0-9]+}", api.PutCharacterHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/character/{characterId:[0-9]+}", api.DeleteCharacterHandler).Methods("DELETE", "OPTIONS")

	// Experience
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/experiences", api.GetExperiencesHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/experience", api.PostExperienceHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/experience/{experienceId:[0-9]+}", api.PutExperienceHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/experience/{experienceId:[0-9]+}", api.DeleteExperienceHandler).Methods("DELETE", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/experience/clear-finalized", api.DeleteExperienceClearFinalizedHandler).Methods("DELETE", "OPTIONS")

	// GameDay routes
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday", api.GetGameDaysHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday", api.PostGameDayHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/gameday/initialize", api.PostInitializeGameDayHandler).Methods("POST", "OPTIONS")

	// Items
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/items", api.GetItemsHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/item", api.PostItemHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/item/{itemId:[0-9]+}", api.PutItemHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/item/{itemId:[0-9]+}", api.DeleteItemHandler).Methods("DELETE", "OPTIONS")

	// Locations
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/locations", api.GetLocationsHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/location", api.PostLocationHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/location/{locationId:[0-9]+}", api.PutLocationHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/location/{locationId:[0-9]+}", api.DeleteLocationHandler).Methods("DELETE", "OPTIONS")

	// Note routes
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{referenceType}", api.GetNotesHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/note", api.PostNoteHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{noteId:[0-9]+}", api.PutNoteHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/note/{noteId:[0-9]+}", api.DeleteNoteHandler).Methods("DELETE", "OPTIONS")

	// Points of interest
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/points-of-interest", api.GetPointsOfInterestHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/point-of-interest", api.PostPointOfInterestHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/point-of-interest/{pointOfInterestId:[0-9]+}", api.PutPointOfInterestHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/point-of-interest/{pointOfInterestId:[0-9]+}", api.DeletePointOfInterestHandler).Methods("DELETE", "OPTIONS")

	// Spell Books
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbooks", api.GetSpellBooksHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook", api.PostSpellBookHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}", api.GetSpellBookHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}", api.DeleteSpellBookHandler).Methods("DELETE", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}", api.PutSpellBookHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}/spell", api.PostSpellBookEntryHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/spellbook/{spellBookId:[0-9]+}/spell/{spellBookEntryId}", api.DeleteSpellBookEntryHandler).Methods("DELETE", "OPTIONS")

	// Spells
	authenticated.HandleFunc("/spells", api.ListSpellsHandler).Methods("GET")

	// Sublocations
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/sublocations", api.GetSublocationsHandler).Methods("GET")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/sublocation", api.PostSublocationHandler).Methods("POST", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/sublocation/{sublocationId:[0-9]+}", api.PutSublocationHandler).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/campaign/{campaignId:[0-9]+}/sublocation/{sublocationId:[0-9]+}", api.DeleteSublocationHandler).Methods("DELETE", "OPTIONS")

	// Start the HTTP server
	log.Println("Server listening on :" + config.Cfg().Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Cfg().Server.Port, r))
}
