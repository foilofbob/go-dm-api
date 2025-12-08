package domain

type Category struct {
	ID         int
	CampaignID int
	Name       string
}

func GetCategory(id int) (*Category, error) {
	row := GetByID("category", id)
	category := &Category{}
	readErr := row.Scan(&category.ID, &category.CampaignID, &category.Name)
	if readErr != nil {
		return nil, readErr
	}
	return category, nil
}

func GetCategories(campaignID int) ([]Category, error) {
	query := "SELECT * FROM category WHERE campaign_id = ?"
	rows, err := DBQuery(query, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category

		if err := rows.Scan(&category.ID, &category.CampaignID, &category.Name); err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return categories, err
	}

	return categories, nil
}

func CreateCategory(campaignID int, name string) (*Category, error) {
	db := DBConnection()
	defer db.Close()

	query := "INSERT INTO category (campaign_id, name) VALUES (?, ?)"
	res, insertErr := db.Exec(query, campaignID, name)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	return GetCategory(int(lid))
}

func UpdateCategory(categoryID int, name string) (*Category, error) {
	db := DBConnection()
	defer db.Close()

	query := "UPDATE category SET name = ? WHERE id = ?"
	_, updateErr := db.Exec(query, name, categoryID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetCategory(categoryID)
}

func DeleteCategory(categoryID int) error {
	db := DBConnection()
	defer db.Close()

	// TODO: Delete associated notes

	query := "DELETE FROM category WHERE id = ?"
	_, deleteErr := db.Exec(query, categoryID)
	if deleteErr != nil {
		println("Delete category error: " + deleteErr.Error())
	}
	return deleteErr
}
