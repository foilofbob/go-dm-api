package api

import (
	"fmt"
	"go_dm_api/domain"
	"go_dm_api/utils"
	"net/http"
)

func ListSpellsHandler(w http.ResponseWriter, _ *http.Request) {
	spells, err := domain.ListSpells()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching spells: %s", err.Error()), http.StatusNotFound)
		return
	}

	utils.StandardResponse(w, spells)
}
