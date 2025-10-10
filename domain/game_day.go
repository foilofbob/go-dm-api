package domain

type GameDay struct {
	ID         int
	CampaignID int
	InGameDay  int
	Day        int
	Month      int
	Year       int
}

type SkinnyCycle struct {
	ID     int
	Offset int
}
type InitializeGameDay struct {
	GameDay GameDay
	Cycles  []SkinnyCycle
}

func GetGameDay(id int) (*GameDay, error) {
	row := GetByID("game_day", id)
	gameDay := &GameDay{}
	readErr := row.Scan(&gameDay.ID, &gameDay.CampaignID, &gameDay.InGameDay, &gameDay.Day, &gameDay.Month, &gameDay.Year)
	if readErr != nil {
		return nil, readErr
	}
	return gameDay, nil
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

func GetMostRecentGameDay(campaignId int) (*GameDay, error) {
	db := DBConnection()
	defer db.Close()

	query := "SELECT * FROM game_day WHERE campaign_id = ? ORDER BY in_game_day DESC limit 1"
	row := db.QueryRow(query, campaignId)
	gameDay := &GameDay{}

	readErr := row.Scan(&gameDay.ID, &gameDay.CampaignID, &gameDay.InGameDay, &gameDay.Day, &gameDay.Month, &gameDay.Year)
	if readErr != nil {
		return nil, readErr
	}
	return gameDay, nil
}

func CreateGameDay(campaignID int, inGameDay int, day int, month int, year int) (*GameDay, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO game_day (campaign_id, in_game_day, day, month, year) VALUES (?, ?, ?, ?, ?)"
	res, insertErr := db.Exec(query, campaignID, inGameDay, day, month, year)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, insertErr
	}

	return GetGameDay(int(lid))
}
