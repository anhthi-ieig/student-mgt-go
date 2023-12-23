package mdw

import (
	"context"
	"net/http"
	"student-service/pkg/service"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuthWithUserService(service *service.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			username, password, ok := c.Request().BasicAuth()

			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing username, password")
			}

			user, err := service.GetUserByUserName(context.Background(), username)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "something when wrong!")
			}
			if user.Username != username {
				return echo.NewHTTPError(http.StatusUnauthorized, "username/password not correct")
			}
			passwordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

			if passwordError != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "username/password not correct")
			}

			c.Set("username", username)
			return next(c)
		}
	}
}
