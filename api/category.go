package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_dm_api/domain"
	"net/http"
	"strconv"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignIdStr := vars["campaignId"]

	campaignId, err := strconv.Atoi(campaignIdStr)

	categories, err := domain.GetCategories(campaignId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Categories fetch failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	StandardResponse(w, categories)
}

func PostCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var category domain.Category
	json.NewDecoder(r.Body).Decode(&category)

	newCategory, err := domain.CreateCategory(category.CampaignID, category.Name)
	if err != nil {
		http.Error(w, "Failed to create category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, newCategory)
}

func PutCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	var category domain.Category
	json.NewDecoder(r.Body).Decode(&category)

	updatedCategory, err := domain.UpdateCategory(category.ID, category.Name)
	if err != nil {
		http.Error(w, "Failed to update category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	StandardResponse(w, updatedCategory)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		StandardResponse(w, nil)
		return
	}
	enableCORS(&w)

	vars := mux.Vars(r)
	idStr := vars["categoryId"]
	categoryID, _ := strconv.Atoi(idStr)

	err := domain.DeleteCategory(categoryID)
	if err != nil {
		http.Error(w, "Failed to delete category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
