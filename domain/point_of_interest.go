package domain

type PointOfInterest struct {
	ID            int
	CampaignID    int
	SublocationID int
	Name          string
}

func GetPointOfInterest(id int) (*PointOfInterest, error) {
	row := GetByID("point_of_interest", id)
	pointOfInterest := &PointOfInterest{}
	readErr := row.Scan(&pointOfInterest.ID, &pointOfInterest.CampaignID, &pointOfInterest.SublocationID, &pointOfInterest.Name)
	if readErr != nil {
		return nil, readErr
	}
	return pointOfInterest, nil
}

func GetPointsOfInterest(campaignID int) ([]PointOfInterest, error) {
	query := "SELECT * FROM point_of_interest WHERE campaign_id = ?"
	rows, err := DBQuery(query, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pointOfInterests []PointOfInterest
	for rows.Next() {
		var pointOfInterest PointOfInterest

		if err := rows.Scan(&pointOfInterest.ID, &pointOfInterest.CampaignID, &pointOfInterest.SublocationID, &pointOfInterest.Name); err != nil {
			return pointOfInterests, err
		}

		pointOfInterests = append(pointOfInterests, pointOfInterest)
	}

	if err = rows.Err(); err != nil {
		return pointOfInterests, err
	}

	return pointOfInterests, nil
}

func CreatePointOfInterest(campaignID int, sublocationID int, name string) (*PointOfInterest, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO point_of_interest (campaign_id, sublocation_id, name) VALUES (?, ?, ?)"
	res, insertErr := db.Exec(query, campaignID, sublocationID, name)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	return GetPointOfInterest(int(lid))
}

func UpdatePointOfInterest(pointOfInterestID int, name string) (*PointOfInterest, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE point_of_interest SET name = ? WHERE id = ?"
	_, updateErr := db.Exec(query, name, pointOfInterestID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetPointOfInterest(pointOfInterestID)
}

func DeletePointOfInterest(pointOfInterestID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM point_of_interest WHERE id = ?"
	_, deleteErr := db.Exec(query, pointOfInterestID)

	// TODO: Delete associated notes

	if deleteErr != nil {
		println("Delete point of interest error: " + deleteErr.Error())
	}

	return deleteErr
}
