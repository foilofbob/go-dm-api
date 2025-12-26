package api

import (
	"fmt"
	"go_dm_api/domain"
	"net/http"
)

func ListSpellsHandler(w http.ResponseWriter, _ *http.Request) {
	spells, err := domain.ListSpells()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching spells: %s", err.Error()), http.StatusNotFound)
		return
	}

	StandardResponse(w, spells)
}
