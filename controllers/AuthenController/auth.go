package AuthenController

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

const jwtSecretKey = "progrise-key" //secret

const tmpUser = "hiden"
const tmpPass = "$2y$12$RPtTxBxB1mH8x..g94LOh.wWeRzHR8Fx0BytDBsEmLt4FCpBqw6ya"

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	match := CheckPasswordHash(password, tmpPass)
	fmt.Println("Match:   ", match)

	// Throws unauthorized error
	if username != tmpUser || match != true {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims

	userName := "Hiden Aruto"
	authority := "President"
	expirationTime := time.Now().Add(time.Minute * 1).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userName
	claims["admin"] = authority
	claims["exp"] = expirationTime

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
