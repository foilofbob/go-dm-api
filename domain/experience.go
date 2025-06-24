package domain

type Experience struct {
	ID          int
	CampaignID  int
	Description string
	XP          int
	Finalized   bool
}

func GetExperience(id int) (*Experience, error) {
	row := GetByID("experience", id)
	experience := &Experience{}
	readErr := row.Scan(&experience.ID, &experience.CampaignID, &experience.Description, &experience.XP, &experience.Finalized)
	if readErr != nil {
		return nil, readErr
	}
	return experience, nil
}

func GetExperiences(campaignID int) ([]Experience, error) {
	query := "SELECT * FROM experience WHERE campaign_id = ?"
	rows, err := DBQuery(query, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []Experience
	for rows.Next() {
		var experience Experience

		if err := rows.Scan(&experience.ID, &experience.CampaignID, &experience.Description, &experience.XP, &experience.Finalized); err != nil {
			return experiences, err
		}

		experiences = append(experiences, experience)
	}

	if err = rows.Err(); err != nil {
		return experiences, err
	}

	return experiences, nil
}

func CreateExperience(campaignID int, description string, xp int, finalized bool) (*Experience, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO experience (campaign_id, description, xp, finalized) VALUES (?, ?, ?, ?)"
	res, insertErr := db.Exec(query, campaignID, description, xp, finalized)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, insertErr
	}

	return GetExperience(int(lid))
}

func UpdateExperience(experienceID int, description string, xp int, finalized bool) (*Experience, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE experience SET description = ?, xp = ?, finalized = ? WHERE id = ?"
	_, updateErr := db.Exec(query, description, xp, finalized, experienceID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetExperience(experienceID)
}

func DeleteExperience(experienceID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM experience WHERE id = ?"
	res, deleteErr := db.Exec(query, experienceID)

	if deleteErr != nil {
		println("Delete experience error: " + deleteErr.Error())
	}

	println(res.RowsAffected())

	return deleteErr
}
