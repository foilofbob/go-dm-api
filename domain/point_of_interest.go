package domain

type PointOfInterest struct {
	ID            int
	SublocationID int
	Name          string
}

func GetPointOfInterest(id int) (*PointOfInterest, error) {
	row := GetByID("point_of_interests", id)
	pointOfInterest := &PointOfInterest{}
	readErr := row.Scan(&pointOfInterest.ID, &pointOfInterest.SublocationID, &pointOfInterest.Name)
	if readErr != nil {
		return nil, readErr
	}
	return pointOfInterest, nil
}

func GetPointsOfInterest(sublocationID int) ([]PointOfInterest, error) {
	query := "SELECT * FROM point_of_interest WHERE sublocation_id = ?"
	rows, err := DBQuery(query, sublocationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pointOfInterests []PointOfInterest
	for rows.Next() {
		var pointOfInterest PointOfInterest

		if err := rows.Scan(&pointOfInterest.ID, &pointOfInterest.SublocationID, &pointOfInterest.Name); err != nil {
			return pointOfInterests, err
		}

		pointOfInterests = append(pointOfInterests, pointOfInterest)
	}

	if err = rows.Err(); err != nil {
		return pointOfInterests, err
	}

	return pointOfInterests, nil
}

func CreatePointOfInterest(locationID int, name string) (*PointOfInterest, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO point_of_interest (sublocation_id, name) VALUES (?, ?)"
	res, insertErr := db.Exec(query, locationID, name)
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

	if deleteErr != nil {
		println("Delete point of interest error: " + deleteErr.Error())
	}

	return deleteErr
}
