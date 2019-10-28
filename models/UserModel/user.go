package UserModel

type User struct {
	key      string
	ID       string `db:"ID"`
	UserName string `db:"USER_NAME"`
}

func AllUsers() (interface{}, error) {

	// UserResult := []User{}
	// db.Select(&UserResult, "SELECT ID,USER_NAME FROM SUB_MST_USER")

	UserResult := []User{
		{ID: "03", UserName: "Horobi"},
		{ID: "04", UserName: "Jin"},
		{ID: "05", UserName: "Valkyrie"},
	}

	// initKey := 0

	// for i := range UserResult {
	// 	initKey++
	// 	UserResult[i].key = strconv.Itoa(initKey)
	// }

	// fmt.Println(UserResult)

	return UserResult, nil
}
