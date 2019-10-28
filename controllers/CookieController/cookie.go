package CookieController

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func WriteCookie(c echo.Context) error {

	name := c.FormValue("username")
	token := c.FormValue("token")
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = time.Now().Add(10 * time.Minute)
	cookie.Path = "/"
	cookie.HttpOnly = false
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie of: "+name)
}

func ReadCookie(c echo.Context) error {
	name := c.FormValue("username")
	cookie, err := c.Cookie(name)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusNotFound, err.Error())
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie: "+cookie.Value)
}

func ReadAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all cookie")
}
