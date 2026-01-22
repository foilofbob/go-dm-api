package domain

type SpellBookEntry struct {
	ID          int
	SpellBookID int
	SpellID     int
}

func GetSpellBookEntry(id int) (*SpellBookEntry, error) {
	row := GetByID("spell_book_entry", id)
	spellBookEntry := &SpellBookEntry{}
	readErr := row.Scan(&spellBookEntry.ID, &spellBookEntry.SpellBookID, &spellBookEntry.SpellID)
	if readErr != nil {
		return nil, readErr
	}
	return spellBookEntry, nil
}

func GetSpellBookEntries(spellBookID int) ([]SpellBookEntry, error) {
	query := "SELECT * FROM spell_book_entry WHERE spell_book_id = ?"
	rows, err := DBQuery(query, spellBookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spellBookEntries []SpellBookEntry
	for rows.Next() {
		var spellBookEntry SpellBookEntry

		if err := rows.Scan(&spellBookEntry.ID, &spellBookEntry.SpellBookID, &spellBookEntry.SpellID); err != nil {
			return spellBookEntries, err
		}

		spellBookEntries = append(spellBookEntries, spellBookEntry)
	}

	if err = rows.Err(); err != nil {
		return spellBookEntries, err
	}

	return spellBookEntries, nil
}

func CreateSpellBookEntry(spellBookID int, spellID int) (*SpellBookEntry, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO spell_book_entry (spell_book_id, spell_id) VALUES (?, ?)"
	res, insertErr := db.Exec(query, spellBookID, spellID)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, insertErr
	}

	return GetSpellBookEntry(int(lid))
}

func DeleteSpellBookEntry(spellBookEntryID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM spell_book_entry WHERE id = ?"
	res, deleteErr := db.Exec(query, spellBookEntryID)

	if deleteErr != nil {
		println("Delete spell book entry error: " + deleteErr.Error())
	}

	println(res.RowsAffected())

	return deleteErr
}
