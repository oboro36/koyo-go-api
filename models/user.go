package models

type User struct {
	ID       string `db:"ID"`
	UserName string `db:"USER_NAME"`
}

func AllUsers() (interface{}, error) {

	UserResult := []User{}
	db.Select(&UserResult, "SELECT ID,USER_NAME FROM SUB_MST_USER")

	return UserResult, nil
}
