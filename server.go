package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/oboro36/koyo/models"
)

func main() {
	//Echo object
	e := echo.New()

	//---------------------Echo Middleware---------------------
	e.Use(middleware.CORS())                                   //Enable cross-domain transfer
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //Server log with content filtering
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.Use(middleware.Recover()) //Recover from panic
	//[END]----------------Echo Middleware---------------------

	//---------------------Route---------------------
	e.POST("/alluser", func(c echo.Context) error {
		models.InitDB("sqlserver", "server=192.168.202.1;user id=collect;password=collect;db=mpacs;")
		res, err := models.AllUsers()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, res)
	})
	//[END]----------------Route---------------------

	e.Logger.Fatal(e.Start(":1323")) //Run on port :1323
}
