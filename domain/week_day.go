package domain

type WeekDay struct {
	ID                int
	CampaignSettingID int
	Name              string
	Order             int
}

func GetWeekDays(campaignSettingID int) ([]WeekDay, error) {
	query := "SELECT * FROM week_day WHERE campaign_setting_id = ?"
	rows, err := DBQuery(query, campaignSettingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weekDays []WeekDay
	for rows.Next() {
		var weekDay WeekDay

		if err := rows.Scan(&weekDay.ID, &weekDay.CampaignSettingID, &weekDay.Name, &weekDay.Order); err != nil {
			return weekDays, err
		}

		weekDays = append(weekDays, weekDay)
	}

	if err = rows.Err(); err != nil {
		return weekDays, err
	}

	return weekDays, nil
}
