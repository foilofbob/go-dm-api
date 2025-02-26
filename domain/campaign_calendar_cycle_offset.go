package domain

type CampaignCalendarCycleOffset struct {
	ID              int
	CampaignID      int
	CalendarCycleID int
	Offset          int
}

func GetCalendarCycleOffsets(campaignID int) ([]CampaignCalendarCycleOffset, error) {
	query := "SELECT * FROM campaign_calendar_cycle_offset WHERE campaign_id = ?"
	rows, err := DBQuery(query, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offsets []CampaignCalendarCycleOffset

	for rows.Next() {
		var offset CampaignCalendarCycleOffset

		if err := rows.Scan(&offset.ID, &offset.CampaignID, &offset.CalendarCycleID, &offset.Offset); err != nil {
			return offsets, err
		}

		offsets = append(offsets, offset)
	}

	if err = rows.Err(); err != nil {
		return offsets, err
	}

	return offsets, nil
}
