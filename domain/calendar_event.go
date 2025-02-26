package domain

type CalendarEvent struct {
	ID                int
	CampaignSettingID int
	Month             int
	Day               int
	Name              string
	Description       string
}

func GetCalendarEvents(campaignSettingID int) ([]CalendarEvent, error) {
	query := "SELECT * FROM calendar_event WHERE campaign_setting_id = ?"
	rows, err := DBQuery(query, campaignSettingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []CalendarEvent

	for rows.Next() {
		var event CalendarEvent

		if err := rows.Scan(&event.ID, &event.CampaignSettingID, &event.Month, &event.Day, &event.Name, &event.Description); err != nil {
			return events, err
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return events, err
	}

	return events, nil
}
