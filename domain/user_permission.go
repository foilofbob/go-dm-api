package domain

import "fmt"

type UserPermission struct {
	ID         int
	UserID     int
	CampaignID int
}

func GetUserPermission(userId int, campaignId int) (*UserPermission, error) {
	db := DBConnection()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM `user_permission` WHERE user_id = ? AND campaign_id = ?", userId, campaignId)
	userPermission := &UserPermission{}
	readErr := row.Scan(&userPermission.ID, &userPermission.UserID, &userPermission.CampaignID)
	if readErr != nil {
		fmt.Println("GetUserPermission err: ", readErr)
		return nil, readErr
	}
	return userPermission, nil
}

func HasPermission(userId int, campaignId int) bool {
	userPermission, err := GetUserPermission(userId, campaignId)

	if err != nil || userPermission == nil {
		return false
	}
	return true
}
