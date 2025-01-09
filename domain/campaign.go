package domain

type Campaign struct {
	ID              int
	Name            string
	CurrentPlayerXP int
}

func GetCampaign(id int) (*Campaign, error) {
	row := GetByID("campaign", id)
	campaign := &Campaign{}
	readErr := row.Scan(&campaign.ID, &campaign.Name, &campaign.CurrentPlayerXP)
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

		if err := rows.Scan(&campaign.ID, &campaign.Name, &campaign.CurrentPlayerXP); err != nil {
			return campaigns, err
		}

		campaigns = append(campaigns, campaign)
	}

	if err = rows.Err(); err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func CreateCampaign(name string) error {
	query := "INSERT INTO campaign (name, current_player_xp) VALUES (?, 0)"
	_, insertErr := DBExec(query, name)
	if insertErr != nil {
		return insertErr
	}
	return nil
}
