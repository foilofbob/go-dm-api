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
	Misc              string
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

		if err := rows.Scan(&player.ID, &player.CampaignID, &player.Name, &player.Race, &player.Class, &player.ArmorClass, &player.HitPoints, &player.PassivePerception, &player.Languages, &player.Movement, &player.Misc); err != nil {
			return players, err
		}

		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return players, err
	}

	return players, nil
}

// TODO: Add Player
