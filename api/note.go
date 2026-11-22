package api

import (
	"encoding/json"
	"fmt"
	"go_dm_api/domain"
	"go_dm_api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetNoteHandler Do we actually need to fetch a single note?
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["noteId"]

	noteId, err := strconv.Atoi(idStr)

	note, err := domain.GetNote(noteId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Note fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	utils.StandardResponse(w, note)
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

	utils.StandardResponse(w, notes)
}

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.StandardResponse(w, nil)
		return
	}
	utils.EnableCORS(&w)

	var note domain.Note
	json.NewDecoder(r.Body).Decode(&note)

	newNote, err := domain.CreateNote(note.CampaignID, note.ReferenceType, note.ReferenceID, note.Category, note.Title, note.Content)
	if err != nil {
		http.Error(w, "Failed to create note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.StandardResponse(w, newNote)
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.StandardResponse(w, nil)
		return
	}
	utils.EnableCORS(&w)

	var note domain.Note
	json.NewDecoder(r.Body).Decode(&note)

	updatedNote, err := domain.UpdateNote(note.ID, note.Category.String, note.Title, note.Content)
	if err != nil {
		http.Error(w, "Failed to update note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.StandardResponse(w, updatedNote)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.StandardResponse(w, nil)
		return
	}
	utils.EnableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["noteId"]
	noteID, _ := strconv.Atoi(idStr)

	err := domain.DeleteNote(noteID)
	if err != nil {
		http.Error(w, "Failed to delete note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
