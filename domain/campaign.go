package domain

import (
	"math"
)

type Campaign struct {
	ID                int
	Name              string
	CampaignSettingID int
	CurrentPlayerXP   int
}

func GetCampaign(id int) (*Campaign, error) {
	row := GetByID("campaign", id)
	campaign := &Campaign{}
	readErr := row.Scan(&campaign.ID, &campaign.Name, &campaign.CurrentPlayerXP, &campaign.CampaignSettingID)
	if readErr != nil {
		return nil, readErr
	}
	return campaign, nil
}

func ListCampaigns() ([]Campaign, error) {
	rows, err := GetAll("campaign")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var campaign Campaign

		if err := rows.Scan(&campaign.ID, &campaign.Name, &campaign.CurrentPlayerXP, &campaign.CampaignSettingID); err != nil {
			return campaigns, err
		}

		campaigns = append(campaigns, campaign)
	}

	if err = rows.Err(); err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func CreateCampaign(name string, currentXp int, campaignSettingID int) (*Campaign, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO campaign (name, current_player_xp, campaign_setting_id) VALUES (?,?,?)"
	res, insertErr := db.Exec(query, name, currentXp, campaignSettingID)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetCampaign(int(lid))
}

func UpdateCurrentPlayerXP(campaignID int, experience int) (int, error) {
	campaign, err := GetCampaign(campaignID)
	if err != nil {
		return -1, err
	}

	if campaign == nil {
		return -1, err
	}

	players, err := GetPlayerCharacters(campaignID)
	if err != nil {
		return -1, err
	}

	avgXP := float64(experience)
	if playerCount := len(players); playerCount > 1 {
		avgXP = math.Round(avgXP / float64(playerCount))
	}
	newXP := campaign.CurrentPlayerXP + int(avgXP)

	query := "UPDATE campaign SET current_player_xp = ? where id = ?"
	_, updateErr := DBExec(query, newXP, campaignID)
	if updateErr != nil {
		return -1, updateErr
	}

	return newXP, nil
}
