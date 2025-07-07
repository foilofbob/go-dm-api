package domain

type Location struct {
	ID         int
	CampaignID int
	Name       string
}

func GetLocation(id int) (*Location, error) {
	row := GetByID("location", id)
	location := &Location{}
	readErr := row.Scan(&location.ID, &location.CampaignID, &location.Name)
	if readErr != nil {
		return nil, readErr
	}
	return location, nil
}

func GetLocations(campaignID int) ([]Location, error) {
	query := "SELECT * FROM location WHERE campaign_id = ?"
	rows, err := DBQuery(query, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var location Location

		if err := rows.Scan(&location.ID, &location.CampaignID, &location.Name); err != nil {
			return locations, err
		}

		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return locations, err
	}

	return locations, nil
}

func CreateLocation(campaignID int, name string) (*Location, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO location (campaign_id, name) VALUES (?, ?)"
	res, insertErr := db.Exec(query, campaignID, name)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	return GetLocation(int(lid))
}

func UpdateLocation(locationID int, name string) (*Location, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE location SET name = ? WHERE id = ?"
	_, updateErr := db.Exec(query, name, locationID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetLocation(locationID)
}

func DeleteLocation(locationID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM location WHERE id = ?"
	_, deleteErr := db.Exec(query, locationID)
	if deleteErr != nil {
		println("Delete location error: " + deleteErr.Error())
	}
	return deleteErr
}
