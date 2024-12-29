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

	// Define routes
	r.HandleFunc("/campaign/{id}", api.GetCampaignHandler).Methods("GET")
	r.HandleFunc("/campaign", api.ListCampaignHandler).Methods("GET")

	// Start the HTTP server
	log.Println("Server listening on :" + config.Cfg().Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Cfg().Server.Port, r))
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
