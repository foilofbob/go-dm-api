package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go_dm_api/config"
	"log"
	"net/http"
	"strconv"
)

const (
	dbDriver = "mysql"
	dbName   = "dm_campaign_manager"
)

func main() {
	// Initialize config
	config.InitConfig()

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	//r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/campaign/{id}", getCampaignHandler).Methods("GET")
	//r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	//r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	// Start the HTTP server
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}

//func createUserHandler(w http.ResponseWriter, r *http.Request) {
//	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
//	if err != nil {
//		panic(err.Error())
//	}
//	defer db.Close()
//
//	// Parse JSON data from the request body
//	var user User
//	json.NewDecoder(r.Body).Decode(&user)
//
//	CreateUser(db, user.Name, user.Email)
//	if err != nil {
//		http.Error(w, "Failed to create user", http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//	fmt.Fprintln(w, "User created successfully")
//}
//
//func CreateUser(db *sql.DB, name, email string) error {
//	query := "INSERT INTO users (name, email) VALUES (?, ?)"
//	_, err := db.Exec(query, name, email)
//	if err != nil {
//		return err
//	}
//	return nil
//}

type Campaign struct {
	ID              int
	Name            string
	CurrentPlayerXP int
}

func getCampaignHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, config.Cfg().Database.Username+":"+config.Cfg().Database.Password+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	campaignId, err := strconv.Atoi(idStr)

	log.Println(fmt.Sprintf("Fetching campaign: %s", campaignId))

	// Call the GetUser function to fetch the user data from the database
	user, err := GetCampaign(db, campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Campaign not found: %s", err.Error()), http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetCampaign(db *sql.DB, id int) (*Campaign, error) {
	query := "SELECT * FROM campaign WHERE id = ?"
	row := db.QueryRow(query, id)

	campaign := &Campaign{}
	err := row.Scan(&campaign.ID, &campaign.Name, &campaign.CurrentPlayerXP)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
