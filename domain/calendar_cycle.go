package domain

type CalendarCycle struct {
	ID                int
	CampaignSettingID int
	Name              string
	Period            int
}

func GetCalendarCycles(campaignSettingID int) ([]CalendarCycle, error) {
	query := "SELECT * FROM calendar_cycle WHERE campaign_setting_id = ?"
	rows, err := DBQuery(query, campaignSettingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cycles []CalendarCycle

	for rows.Next() {
		var cycle CalendarCycle

		if err := rows.Scan(&cycle.ID, &cycle.CampaignSettingID, &cycle.Name, &cycle.Period); err != nil {
			return cycles, err
		}

		cycles = append(cycles, cycle)
	}

	if err = rows.Err(); err != nil {
		return cycles, err
	}

	return cycles, nil
}
