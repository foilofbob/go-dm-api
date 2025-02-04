package domain

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

func CreateCampaign(name string, campaignSettingID int) error {
	query := "INSERT INTO campaign (name, current_player_xp, campaign_setting_id) VALUES (?, 0, ?)"
	_, insertErr := DBExec(query, name, campaignSettingID)
	if insertErr != nil {
		return insertErr
	}
	return nil
}
