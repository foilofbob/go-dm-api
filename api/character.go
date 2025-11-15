package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func GetPlayersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	characters, err := domain.GetPlayerCharacters(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Player characters fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, characters)
}

func GetNPCsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	characters, err := domain.GetNonPlayerCharacters(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("NPCs fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, characters)
}

func PostCharacterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var character domain.Character
	json.NewDecoder(r.Body).Decode(&character)

	newCharacter, err := domain.CreateCharacter(character.CampaignID, character.Name, character.Race, character.Class,
		character.ArmorClass, character.HitPoints, character.PassivePerception, character.Languages, character.Movement,
		character.Strength, character.Dexterity, character.Constitution, character.Intelligence, character.Wisdom,
		character.Charisma, character.Proficiencies, character.PlayerType, character.StrengthSaveProficiency,
		character.DexteritySaveProficiency, character.ConstitutionSaveProficiency, character.IntelligenceSaveProficiency,
		character.WisdomSaveProficiency, character.CharismaSaveProficiency, character.AcrobaticsProficiencyBonus,
		character.AnimalHandlingProficiencyBonus, character.ArcanaProficiencyBonus, character.AthleticsProficiencyBonus,
		character.DeceptionProficiencyBonus, character.HistoryProficiencyBonus, character.InsightProficiencyBonus,
		character.IntimidationProficiencyBonus, character.InvestigationProficiencyBonus, character.MedicineProficiencyBonus,
		character.NatureProficiencyBonus, character.PerceptionProficiencyBonus, character.PerformanceProficiencyBonus,
		character.PersuasionProficiencyBonus, character.ReligionProficiencyBonus, character.SleightOfHandProficiencyBonus,
		character.StealthProficiencyBonus, character.SurvivalProficiencyBonus, character.Level)
	if err != nil {
		http.Error(w, "Failed to create character: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newCharacter)
}

func PutCharacterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	fmt.Println("Received PUT Character Request")

	var character domain.Character
	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		http.Error(w, "Failed to update character: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	updateItem, err := domain.UpdateCharacter(character.ID, character.Name, character.Race, character.Class,
		character.ArmorClass, character.HitPoints, character.PassivePerception, character.Languages, character.Movement,
		character.Strength, character.Dexterity, character.Constitution, character.Intelligence, character.Wisdom,
		character.Charisma, character.Proficiencies, character.PlayerType, character.StrengthSaveProficiency,
		character.DexteritySaveProficiency, character.ConstitutionSaveProficiency, character.IntelligenceSaveProficiency,
		character.WisdomSaveProficiency, character.CharismaSaveProficiency, character.AcrobaticsProficiencyBonus,
		character.AnimalHandlingProficiencyBonus, character.ArcanaProficiencyBonus, character.AthleticsProficiencyBonus,
		character.DeceptionProficiencyBonus, character.HistoryProficiencyBonus, character.InsightProficiencyBonus,
		character.IntimidationProficiencyBonus, character.InvestigationProficiencyBonus, character.MedicineProficiencyBonus,
		character.NatureProficiencyBonus, character.PerceptionProficiencyBonus, character.PerformanceProficiencyBonus,
		character.PersuasionProficiencyBonus, character.ReligionProficiencyBonus, character.SleightOfHandProficiencyBonus,
		character.StealthProficiencyBonus, character.SurvivalProficiencyBonus, character.Level)
	if err != nil {
		http.Error(w, "Failed to update character: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Completed PUT Character Request")

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updateItem)
}

func DeleteCharacterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["characterId"]
	characterID, _ := strconv.Atoi(idStr)

	err := domain.DeleteCharacter(characterID)
	if err != nil {
		http.Error(w, "Failed to delete character: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
