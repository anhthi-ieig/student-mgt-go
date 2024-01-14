package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
)

type userAPI struct {
	userService interfaces.UserService
}

func (api *userAPI) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User ID")
	}

	class, err := api.userService.Get(c.Request().Context(), id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
	}

	return c.JSON(http.StatusOK, class)
}

func (api *userAPI) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User ID")
	}

	var class model.User
	if err := c.Bind(&class); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	_, err = api.userService.Get(c.Request().Context(), id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
	}

	response, err := api.userService.Update(c.Request().Context(), id, class)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, response)
}

func NewUserAPI(userService interfaces.UserService) *userAPI {
	return &userAPI{userService}
}
