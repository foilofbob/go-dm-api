package domain

type Sublocation struct {
	ID          int
	LocationID  int
	Name        string
	Description string
}

func GetSublocation(id int) (*Sublocation, error) {
	row := GetByID("sublocation", id)
	sublocation := &Sublocation{}
	readErr := row.Scan(&sublocation.ID, &sublocation.LocationID, &sublocation.Name, &sublocation.Description)
	if readErr != nil {
		return nil, readErr
	}
	return sublocation, nil
}

func GetSublocations(locationId int) ([]Sublocation, error) {
	query := "SELECT * FROM sublocation WHERE location_id = ?"
	rows, err := DBQuery(query, locationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sublocations []Sublocation
	for rows.Next() {
		var sublocation Sublocation

		if err := rows.Scan(&sublocation.ID, &sublocation.LocationID, &sublocation.Name, &sublocation.Description); err != nil {
			return sublocations, err
		}

		sublocations = append(sublocations, sublocation)
	}

	if err := rows.Err(); err != nil {
		return sublocations, err
	}

	return sublocations, nil
}

func CreateSublocation(locationId int, name string, description string) (*Sublocation, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO sublocation (location_id, name, description) VALUES (?, ?, ?)"
	res, insertErr := db.Exec(query, locationId, name, description)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetSublocation(int(lid))
}

func UpdateSublocation(id int, name string, description string) (*Sublocation, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE sublocation SET name = ?, description = ? WHERE id = ?"
	_, updateErr := db.Exec(query, name, description, id)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetSublocation(id)
}

func DeleteSublocation(id int) error {
	db := DBConnection()
	defer db.Close()
	query := "DELETE FROM sublocation WHERE id = ?"
	_, deleteErr := db.Exec(query, id)

	if deleteErr != nil {
		println("Delete sublocation error: " + deleteErr.Error())
	}

	return deleteErr
}
