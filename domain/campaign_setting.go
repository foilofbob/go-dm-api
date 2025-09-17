package domain

type CampaignSetting struct {
	ID   int
	Name string
}

func ListCampaignSettings() ([]CampaignSetting, error) {
	rows, err := GetAll("campaign_setting")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaignSettings []CampaignSetting
	for rows.Next() {
		var campaignSetting CampaignSetting

		if err := rows.Scan(&campaignSetting.ID, &campaignSetting.Name); err != nil {
			return campaignSettings, err
		}

		campaignSettings = append(campaignSettings, campaignSetting)
	}

	if err = rows.Err(); err != nil {
		return campaignSettings, err
	}

	return campaignSettings, nil
}
