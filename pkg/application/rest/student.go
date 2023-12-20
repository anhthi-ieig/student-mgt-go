package rest

import (
	"fmt"
	"net/http"
	"student-service/pkg/application/interfaces"

	"github.com/labstack/echo/v4"
)

type studentAPI struct {
	studentService interfaces.StudentService
}

func (api *studentAPI) List(c echo.Context) error {
	students, err := api.studentService.GetStudents(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, students)
}

func NewStudentAPI(studentService interfaces.StudentService) *studentAPI {
	return &studentAPI{studentService}
}
