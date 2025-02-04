package domain

type Month struct {
	ID                int
	CampaignSettingID int
	Name              string
	NumDays           int
	Order             int
}

type CampaignMonthSummary struct {
	MonthCount       int
	CurrentMonthDays int
}

func GetMonths(campaignSettingID int) ([]Month, error) {
	query := "SELECT * FROM month WHERE campaign_setting_id = ?"
	rows, err := DBQuery(query, campaignSettingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var months []Month
	for rows.Next() {
		var month Month

		if err := rows.Scan(&month.ID, &month.CampaignSettingID, &month.Name, &month.NumDays, &month.Order); err != nil {
			return months, err
		}

		months = append(months, month)
	}

	if err = rows.Err(); err != nil {
		return months, err
	}

	return months, nil
}

func GetCampaignMonthSummary(campaignID int, monthNum int) (*CampaignMonthSummary, error) {
	db := DBConnection()
	defer db.Close()

	query := "select count(*), max(if(m.order = ?, m.num_days, 0)) " +
		"from campaign c " +
		"left join month m on m.campaign_setting_id = c.campaign_setting_id " +
		"where c.id = ?"

	row := db.QueryRow(query, monthNum, campaignID)
	campaignMonthSummary := &CampaignMonthSummary{}

	readErr := row.Scan(&campaignMonthSummary.MonthCount, &campaignMonthSummary.CurrentMonthDays)
	if readErr != nil {
		return nil, readErr
	}
	return campaignMonthSummary, nil
}
