package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

// GetNoteHandler Do we actually need to fetch a single note?
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	noteId, err := strconv.Atoi(idStr)

	note, err := domain.GetNote(noteId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Note fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, note)
}

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]
	referenceType := vars["referenceType"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	notes, err := domain.GetNotes(campaignId, referenceType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Notes fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, notes)
}

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note domain.Note
	json.NewDecoder(r.Body).Decode(&note)

	err := domain.CreateNote(note.CampaignID, note.ReferenceType, note.ReferenceID, note.Category, note.Title, note.Content)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Campaign created successfully")
}
