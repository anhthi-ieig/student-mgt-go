package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.Get("username").(string)

	token := jwt.New(jwt.SigningMethodHS256)

	clams := token.Claims.(jwt.MapClaims)

	clams["username"] = username
	clams["exp"] = time.Now().Add(2 * time.Minute).Unix()

	t, err := token.SignedString([]byte("mysecretkey"))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
