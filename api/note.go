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
	idStr := vars["noteId"]

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
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var note domain.Note
	json.NewDecoder(r.Body).Decode(&note)

	newNote, err := domain.CreateNote(note.CampaignID, note.ReferenceType, note.ReferenceID, note.Category, note.Title, note.Content)
	if err != nil {
		http.Error(w, "Failed to create note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newNote)
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var note domain.Note
	json.NewDecoder(r.Body).Decode(&note)

	updatedNote, err := domain.UpdateNote(note.ID, note.Category.String, note.Title, note.Content)
	if err != nil {
		http.Error(w, "Failed to update note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updatedNote)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

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
