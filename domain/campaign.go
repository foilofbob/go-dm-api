package domain

import (
	"database/sql"
)

type Campaign struct {
	ID              int
	Name            string
	CurrentPlayerXP int
}

func GetCampaign(db *sql.DB, id int) (*Campaign, error) {
	query := "SELECT * FROM campaign WHERE id = ?"
	row := db.QueryRow(query, id)

	campaign := &Campaign{}
	err := row.Scan(&campaign.ID, &campaign.Name, &campaign.CurrentPlayerXP)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func ListCampaigns(db *sql.DB) ([]Campaign, error) {
	query := "SELECT * FROM campaign"
	rows, err := db.Query(query)
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
