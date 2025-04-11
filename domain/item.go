package domain

type Item struct {
	ID           int
	CampaignID   int
	Name         string
	Description  string
	Link         string
	Rarity       string
	Cost         string
	Requirements string
	IsContainer  bool
	CarriedBy    string
	CarriedById  int
}

func GetItems(campaignID int) ([]Item, error) {
	query := "SELECT * FROM item WHERE campaign_id = ?"
	rows, err := DBQuery(query, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item

		if err := rows.Scan(&item.ID, &item.CampaignID, &item.Name, &item.Description, &item.Link, &item.Rarity, &item.Cost, &item.Requirements, &item.IsContainer, &item.CarriedBy, &item.CarriedById); err != nil {
			return items, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}
