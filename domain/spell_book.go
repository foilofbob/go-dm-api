package domain

type SpellBook struct {
	ID          int
	CampaignID  int
	CharacterID int
	SpellStats  string
}

func GetSpellBook(id int) (*SpellBook, error) {
	row := GetByID("spell_book", id)
	spellBook := &SpellBook{}
	readErr := row.Scan(&spellBook.ID, &spellBook.CampaignID, &spellBook.CharacterID, &spellBook.SpellStats)
	if readErr != nil {
		return nil, readErr
	}
	return spellBook, nil
}

func GetSpellBooks(campaignID int) ([]SpellBook, error) {
	query := "SELECT * FROM spell_book WHERE campaign_id = ?"
	rows, err := DBQuery(query, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spellBooks []SpellBook
	for rows.Next() {
		var spellBook SpellBook

		if err := rows.Scan(&spellBook.ID, &spellBook.CampaignID, &spellBook.CharacterID, &spellBook.SpellStats); err != nil {
			return spellBooks, err
		}

		spellBooks = append(spellBooks, spellBook)
	}

	if err = rows.Err(); err != nil {
		return spellBooks, err
	}

	return spellBooks, nil
}

func CreateSpellBook(campaignID int, characterID int, spellStats string) (*SpellBook, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO spell_book (campaign_id, character_id, spell_stats) VALUES (?, ?, ?)"
	res, insertErr := db.Exec(query, campaignID, characterID, spellStats)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, insertErr
	}

	return GetSpellBook(int(lid))
}

func UpdateSpellBook(spellBookID int, spellStats string) (*SpellBook, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE spell_book SET spell_stats = ? WHERE id = ?"
	_, updateErr := db.Exec(query, spellStats, spellBookID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetSpellBook(spellBookID)
}

func DeleteSpellBook(spellBookID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM spell_book WHERE id = ?"
	res, deleteErr := db.Exec(query, spellBookID)

	if deleteErr != nil {
		println("Delete spell book error: " + deleteErr.Error())
	}

	println(res.RowsAffected())

	return deleteErr
}
