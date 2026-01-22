package domain

import (
	"database/sql"
)

type Note struct {
	ID            int
	CampaignID    int
	ReferenceType string
	ReferenceID   int
	Category      sql.NullString
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

func GetNotes(campaignID int, referenceType string) ([]Note, error) {
	query := "SELECT * FROM note WHERE campaign_id = ? AND reference_type = ?"
	rows, err := DBQuery(query, campaignID, referenceType)
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

func CreateNote(campaignID int, referenceType string, referenceID int, category sql.NullString, title string, content string) (*Note, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO note (campaign_id, reference_type, reference_id, category, title, content) VALUES (?, ?, ?, ?, ?, ?)"
	res, insertErr := db.Exec(query, campaignID, referenceType, referenceID, category, title, content)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, insertErr
	}

	return GetNote(int(lid))
}

func UpdateNote(noteID int, category string, title string, content string) (*Note, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE note SET category = ?, title = ?, content = ? WHERE id = ?"
	_, updateErr := db.Exec(query, category, title, content, noteID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetNote(noteID)
}

func DeleteNote(noteID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM note WHERE id = ?"
	res, deleteErr := db.Exec(query, noteID)

	if deleteErr != nil {
		println("Delete note error: " + deleteErr.Error())
	}

	println(res.RowsAffected())

	return deleteErr
}
