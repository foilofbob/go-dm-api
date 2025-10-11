package domain

type Player struct {
	ID                int
	CampaignID        int
	Name              string
	Race              string
	Class             string
	ArmorClass        int
	HitPoints         int
	PassivePerception int
	Languages         string
	Movement          int
	Strength          int
	Dexterity         int
	Constitution      int
	Intelligence      int
	Wisdom            int
	Charisma          int
	Proficiencies     string
	Saves             string
}

func GetPlayer(id int) (*Player, error) {
	row := GetByID("player", id)
	player := &Player{}
	readErr := row.Scan(&player.ID, &player.CampaignID, &player.Name, &player.Race, &player.Class, &player.ArmorClass, &player.HitPoints, &player.PassivePerception,
		&player.Languages, &player.Movement, &player.Strength, &player.Dexterity, &player.Constitution, &player.Intelligence, &player.Wisdom, &player.Charisma, &player.Proficiencies, &player.Saves)
	if readErr != nil {
		return nil, readErr
	}
	return player, nil
}

func GetPlayers(campaignId int) ([]Player, error) {
	query := "SELECT * FROM player WHERE campaign_id = ?"
	rows, err := DBQuery(query, campaignId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var player Player

		if err := rows.Scan(&player.ID, &player.CampaignID, &player.Name, &player.Race, &player.Class, &player.ArmorClass, &player.HitPoints, &player.PassivePerception,
			&player.Languages, &player.Movement, &player.Strength, &player.Dexterity, &player.Constitution, &player.Intelligence, &player.Wisdom, &player.Charisma, &player.Proficiencies, &player.Saves); err != nil {
			return players, err
		}

		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return players, err
	}

	return players, nil
}

func CreatePlayer(campaignID int, name string, race string, class string, armorClass int, hitPoints int, passivePerception int, languages string, movement int,
	strength int, dexterity int,
	constitution int,
	intelligence int,
	wisdom int,
	charisma int,
	proficiencies string,
	saves string) (*Player, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO player (campaign_id, name, race, class, armor_class, hit_points, passive_perception, languages, movement, strength, dexterity, constitution, intelligence, wisdom, charisma, proficiencies, saves) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, insertErr := db.Exec(query, campaignID, name, race, class, armorClass, hitPoints, passivePerception, languages, movement, strength, dexterity, constitution, intelligence, wisdom, charisma, proficiencies, saves)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, insertErr
	}

	return GetPlayer(int(lid))
}

func UpdatePlayer(playerID int, name string, race string, class string, armorClass int, hitPoints int, passivePerception int, languages string, movement int,
	strength int, dexterity int,
	constitution int,
	intelligence int,
	wisdom int,
	charisma int,
	proficiencies string,
	saves string) (*Player, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE player SET name = ?, race = ?, class = ?, armor_class = ?, hit_points = ?, passive_perception = ?, languages = ?, movement = ? " +
		", strength = ? , dexterity = ? , constitution = ? , intelligence = ? , wisdom = ? , charisma = ? , proficiencies = ? , saves = ? WHERE id = ?"
	_, updateErr := db.Exec(query, name, race, class, armorClass, hitPoints, passivePerception, languages, movement, strength, dexterity, constitution, intelligence, wisdom, charisma, proficiencies, saves, playerID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetPlayer(playerID)
}

func DeletePlayer(playerID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM player WHERE id = ?"
	res, deleteErr := db.Exec(query, playerID)

	if deleteErr != nil {
		println("Delete player error: " + deleteErr.Error())
	}

	println(res.RowsAffected())

	return deleteErr
}
