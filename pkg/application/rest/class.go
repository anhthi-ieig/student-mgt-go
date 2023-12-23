package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"

	"github.com/labstack/echo/v4"
)

type classAPI struct {
	classService interfaces.ClassService
}

func (api *classAPI) List(c echo.Context) error {
	classes, err := api.classService.List(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, classes)
}

func (api *classAPI) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	class, err := api.classService.Get(c.Request().Context(), id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "Class Not Found")
	}

	return c.JSON(http.StatusOK, class)
}

func (api *classAPI) Create(c echo.Context) error {
	var class model.Class
	if err := c.Bind(&class); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	response, err := api.classService.Create(c.Request().Context(), class)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusCreated, response)
}

func (api *classAPI) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	var class model.Class
	if err := c.Bind(&class); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	_, err = api.classService.Get(c.Request().Context(), id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "Class Not Found")
	}

	response, err := api.classService.Update(c.Request().Context(), id, class)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, response)
}

func (api *classAPI) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	err = api.classService.Delete(c.Request().Context(), id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (api *classAPI) ListStudents(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	students, err := api.classService.ListStudents(c.Request().Context(), id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, students)
}

func (api *classAPI) ListTeachers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	teachers, err := api.classService.ListTeachers(c.Request().Context(), id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, teachers)
}

func (api *classAPI) AddStudent(c echo.Context) error {
	classId, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	studentId, err := strconv.Atoi(c.Param("student-id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Student ID")
	}

	err = api.classService.AddStudent(c.Request().Context(), classId, studentId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, nil)
}

func (api *classAPI) AddTeacher(c echo.Context) error {
	classId, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	teacherId, err := strconv.Atoi(c.Param("teacher-id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Teacher ID")
	}

	err = api.classService.AddTeacher(c.Request().Context(), classId, teacherId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, nil)
}

func (api *classAPI) RemoveStudent(c echo.Context) error {
	classId, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	studentId, err := strconv.Atoi(c.Param("student-id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Student ID")
	}

	err = api.classService.RemoveStudent(c.Request().Context(), classId, studentId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, nil)
}

func (api *classAPI) RemoveTeacher(c echo.Context) error {
	classId, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Class ID")
	}

	teacherId, err := strconv.Atoi(c.Param("teacher-id"))
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Teacher ID")
	}

	err = api.classService.RemoveTeacher(c.Request().Context(), classId, teacherId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, nil)
}

func NewClassAPI(classService interfaces.ClassService) *classAPI {
	return &classAPI{classService}
}
