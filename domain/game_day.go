package domain

type GameDay struct {
	ID         int
	CampaignID int
	InGameDay  int
	Day        int
	Month      int
	Year       int
}

func GetGameDays(campaignId int) ([]GameDay, error) {
	query := "SELECT * FROM game_day WHERE campaign_id = ? ORDER BY in_game_day DESC"
	rows, err := DBQuery(query, campaignId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gameDays []GameDay
	for rows.Next() {
		var gameDay GameDay

		if err := rows.Scan(&gameDay.ID, &gameDay.CampaignID, &gameDay.InGameDay, &gameDay.Day, &gameDay.Month, &gameDay.Year); err != nil {
			return gameDays, err
		}

		gameDays = append(gameDays, gameDay)
	}

	if err = rows.Err(); err != nil {
		return gameDays, err
	}

	return gameDays, nil

}

func CreateGameDay(campaignID int, inGameDay int, day int, month int, year int) error {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO game_day (campaign_id, in_game_day, day, month, year) VALUES (?, ?, ?, ?, ?)"
	_, insertErr := db.Exec(query, campaignID, inGameDay, day, month, year)
	if insertErr != nil {
		return insertErr
	}
	return nil
}
