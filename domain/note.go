package domain

import "database/sql"

type Note struct {
	ID            int
	CampaignID    int
	ReferenceType string
	ReferenceID   int
	Category      sql.NullString // TODO: marshaller to handle this
	Title         string
	Content       string
}

func GetNote(id int) (*Note, error) {
	row := GetByID("note", id)
	note := &Note{}
	readErr := row.Scan(&note.ID, &note.CampaignID, &note.ReferenceType, &note.ReferenceID, &note.Category, &note.Title, &note.Content)
	if readErr != nil {
		return nil, readErr
	}
	return note, nil
}

func GetNotes(campaignId int, referenceType string) ([]Note, error) {
	query := "SELECT * FROM note WHERE campaign_id = ? AND reference_type = ?"
	rows, err := DBQuery(query, campaignId, referenceType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note

		if err := rows.Scan(&note.ID, &note.CampaignID, &note.ReferenceType, &note.ReferenceID, &note.Category, &note.Title, &note.Content); err != nil {
			return notes, err
		}

		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return notes, err
	}

	return notes, nil

}

func CreateNote(campaignID int, referenceType string, referenceID int, category sql.NullString, title string, content string) error {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO note (campaign_id, reference_type, reference_id, category, title, content) VALUES (?, ?, ?, ?, ?, ?)"
	_, insertErr := db.Exec(query, campaignID, referenceType, referenceID, category, title, content)
	if insertErr != nil {
		return insertErr
	}
	return nil
}
