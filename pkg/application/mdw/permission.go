package mdw

import (
	"context"
	"net/http"
	"strings"
	"student-service/pkg/data-access/dto"
	"student-service/pkg/service"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsValidToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		splitToken := strings.Split(tokenString, "Bearer ")
		if len(splitToken) != 2 {
			return c.String(http.StatusUnauthorized, "Invalid token format")
		}

		tokenString = splitToken[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "unexpected signing method")
			}
			return []byte("mysecretkey"), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
		}
		username, ok := claims["username"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid username in token")
		}

		c.Set("username", username)

		return next(c)
	}
}

func IsValidPermission(service *service.UserService, role dto.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			username := c.Get("username").(string)

			userFromDB, err := service.GetUserByUserName(context.Background(), username)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "something when wrong!")
			}

			if userFromDB.Role == role {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusForbidden, "Insufficient permissions")
		}
	}
}
