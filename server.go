package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/oboro36/koyo/controllers/AuthenController"
	"github.com/oboro36/koyo/controllers/CookieController"
	"github.com/oboro36/koyo/controllers/UserController"
)

//---------------------Custom Middleware---------------------
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:3000")
		c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, "POST, GET, OPTIONS, PUT, DELETE")
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
		return next(c)
	}
}

//[END]----------------Custom Middleware---------------------

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

	//---------------------Custom Middleware---------------------
	e.Use(ServerHeader)
	//[END]----------------Custom Middleware---------------------

	//---------------------Route---------------------
	//Login route
	e.POST("/login", AuthenController.Login)

	// Unauthenticated route
	e.GET("/", AuthenController.Accessible)

	// // // Restricted group
	// r := e.Group("/restricted")
	// r.Use(middleware.JWT([]byte(jwtSecretKey)))
	// r.GET("", AuthenModule.Restricted)

	//Cookie
	e.POST("/writecookie", CookieController.WriteCookie)
	e.POST("/readcookie", CookieController.ReadCookie)
	e.POST("/readallcookies", CookieController.ReadAllCookies)

	//Other
	e.POST("/alluser", UserController.GetAllUsers)
	//[END]----------------Route---------------------

	e.Logger.Fatal(e.Start(":1323")) //Run on port :1323
}
