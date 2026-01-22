package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func PostSpellBookEntryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var spellBookEntry domain.SpellBookEntry
	json.NewDecoder(r.Body).Decode(&spellBookEntry)

	newSpellBookEntry, err := domain.CreateSpellBookEntry(spellBookEntry.SpellBookID, spellBookEntry.SpellID)
	if err != nil {
		http.Error(w, "Failed to create spell book entry: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newSpellBookEntry)
}

func DeleteSpellBookEntryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["spellBookEntryId"]
	spellBookEntryID, _ := strconv.Atoi(idStr)

	err := domain.DeleteSpellBookEntry(spellBookEntryID)
	if err != nil {
		http.Error(w, "Failed to delete spell book entry: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
