package api

import (
	"fmt"
	"go_dm_api/domain"
	"go_dm_api/utils"
	"net/http"
)

func ListCampaignSettingHandler(w http.ResponseWriter, _ *http.Request) {
	campaignSettings, err := domain.ListCampaignSettings()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching campaign settings: %s", err.Error()), http.StatusNotFound)
		return
	}

	utils.StandardResponse(w, campaignSettings)
}
