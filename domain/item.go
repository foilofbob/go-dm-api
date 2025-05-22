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
	CarriedByID  int
}

func GetItem(id int) (*Item, error) {
	row := GetByID("item", id)
	item := &Item{}
	readErr := row.Scan(&item.ID, &item.CampaignID, &item.Name, &item.Description, &item.Link, &item.Rarity, &item.Cost, &item.Requirements, &item.IsContainer, &item.CarriedBy, &item.CarriedByID)
	if readErr != nil {
		return nil, readErr
	}
	return item, nil
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

		if err := rows.Scan(&item.ID, &item.CampaignID, &item.Name, &item.Description, &item.Link, &item.Rarity, &item.Cost, &item.Requirements, &item.IsContainer, &item.CarriedBy, &item.CarriedByID); err != nil {
			return items, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func CreateItem(campaignID int, name string, description string, link string, rarity string, cost string, requirements string, isContainer bool, carriedBy string, carriedByID int) (*Item, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO item (campaign_id, name, description, link, rarity, cost, requirements, is_container, carried_by, carried_by_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, insertErr := db.Exec(query, campaignID, name, description, link, rarity, cost, requirements, isContainer, carriedBy, carriedByID)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, insertErr
	}

	return GetItem(int(lid))
}

func UpdateItem(itemID int, name string, description string, link string, rarity string, cost string, requirements string, isContainer bool, carriedBy string, carriedByID int) (*Item, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE item SET name = ?, description = ?, link = ?, rarity = ?, cost = ?, requirements = ?, is_container = ?, carried_by = ?, carried_by_id = ? WHERE id = ?"
	_, updateErr := db.Exec(query, name, description, link, rarity, cost, requirements, isContainer, carriedBy, carriedByID, itemID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetItem(itemID)
}

func DeleteItem(itemID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM item WHERE id = ?"
	res, deleteErr := db.Exec(query, itemID)

	if deleteErr != nil {
		println("Delete item error: " + deleteErr.Error())
	}

	println(res.RowsAffected())

	return deleteErr
}
