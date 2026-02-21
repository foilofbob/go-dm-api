package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

type CompleteSpellBook struct {
	ID               int
	CampaignID       int
	CharacterID      int
	SpellStats       string
	SpellBookEntries []domain.SpellBookEntry
}

func BuildCompleteSpellBook(spellBook domain.SpellBook) CompleteSpellBook {
	var completeSpellBook CompleteSpellBook
	spellBookEntries, _ := domain.GetSpellBookEntries(spellBook.ID)

	completeSpellBook.ID = spellBook.ID
	completeSpellBook.CampaignID = spellBook.CampaignID
	completeSpellBook.CharacterID = spellBook.CharacterID
	completeSpellBook.SpellStats = spellBook.SpellStats
	completeSpellBook.SpellBookEntries = spellBookEntries

	return completeSpellBook
}

func GetSpellBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["spellBookId"]

	spellBookId, err := strconv.Atoi(idStr)

	spellBook, err := domain.GetSpellBook(spellBookId)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, fmt.Sprintf("Spell Book fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, BuildCompleteSpellBook(*spellBook))
}

func GetSpellBooksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	spellBooks, err := domain.GetSpellBooks(campaignId)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, fmt.Sprintf("Spell Books fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var completeSpellBooks []CompleteSpellBook
	for _, spellBook := range spellBooks {
		completeSpellBooks = append(completeSpellBooks, BuildCompleteSpellBook(spellBook))
	}

	StandardResponse(w, completeSpellBooks)
}

func PostSpellBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var spellBook domain.SpellBook
	json.NewDecoder(r.Body).Decode(&spellBook)

	newSpellBook, err := domain.CreateSpellBook(spellBook.CampaignID, spellBook.CharacterID, spellBook.SpellStats)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Failed to create spell book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newSpellBook)
}

func PutSpellBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var spellBook domain.SpellBook
	json.NewDecoder(r.Body).Decode(&spellBook)

	updatedSpellBook, err := domain.UpdateSpellBook(spellBook.ID, spellBook.SpellStats)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Failed to update spell book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updatedSpellBook)
}

func DeleteSpellBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["spellBookId"]
	spellBookID, _ := strconv.Atoi(idStr)

	err := domain.DeleteSpellBook(spellBookID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Failed to delete spell book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
