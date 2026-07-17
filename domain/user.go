package domain

type User struct {
	ID       int
	Username string
	Password string
}

func GetUser(id int) (*User, error) {
	row := GetByID("user", id)
	user := &User{}
	readErr := row.Scan(&user.ID, &user.Username, &user.Password)
	if readErr != nil {
		return nil, readErr
	}
	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	db := DBConnection()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM `user` WHERE username = ?", username)
	user := &User{}
	readErr := row.Scan(&user.ID, &user.Username, &user.Password)
	if readErr != nil {
		return nil, readErr
	}
	return user, nil
}
