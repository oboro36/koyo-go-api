package UserController

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/oboro36/koyo/models/UserModel"
)

func GetAllUsers(c echo.Context) error {
	// models.InitDB("sqlserver", "server=192.168.202.1;user id=collect;password=collect;db=mpacs;")
	res, err := UserModel.AllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}
