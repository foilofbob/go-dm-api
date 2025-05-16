package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	items, err := domain.GetItems(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Items fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, items)
}

func PostItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var item domain.Item
	json.NewDecoder(r.Body).Decode(&item)

	newItem, err := domain.CreateItem(item.CampaignID, item.Name, item.Description, item.Link, item.Rarity, item.Cost, item.Requirements, item.IsContainer, item.CarriedBy, item.CarriedByID)
	if err != nil {
		http.Error(w, "Failed to create item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newItem)
}

func PutItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var item domain.Item
	json.NewDecoder(r.Body).Decode(&item)

	updateItem, err := domain.UpdateItem(item.ID, item.Name, item.Description, item.Link, item.Rarity, item.Cost, item.Requirements, item.IsContainer, item.CarriedBy, item.CarriedByID)
	if err != nil {
		http.Error(w, "Failed to update item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updateItem)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["itemId"]
	itemID, _ := strconv.Atoi(idStr)

	err := domain.DeleteItem(itemID)
	if err != nil {
		http.Error(w, "Failed to delete item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
