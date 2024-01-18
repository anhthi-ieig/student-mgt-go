package main

import (
	"net/http"
	"student-service/pkg/application/handler"
	"student-service/pkg/application/mdw"
	"student-service/pkg/application/rest"
	dataaccess "student-service/pkg/data-access"
	"student-service/pkg/data-access/dto"
	"student-service/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	sqlDB := dataaccess.InitializeSequelDB("postgres://user:password@localhost:5432/student-service?sslmode=disable")

	// Data access
	userDA := dataaccess.NewUserDA(sqlDB)
	classDA := dataaccess.NewClassDA(sqlDB)

	// Services
	userService := service.NewUserService(userDA)
	classService := service.NewClassService(classDA)

	// APIs
	userAPI := rest.NewUserAPI(userService)
	classAPI := rest.NewClassAPI(classService)

	server := initializeHTTPServer()

	/*
	 * Routing
	 */

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	server.POST("/login", handler.Login, mdw.BasicAuthWithUserService(userService))

	// Users
	users := server.Group("/users", mdw.IsValidToken)
	users.GET("", userAPI.List)
	users.GET("/:id", userAPI.Get, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher, dto.Role_Student}))
	users.POST("", userAPI.Create, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))
	users.PUT("/:id", userAPI.Update, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin}))

	// Classes
	classes := server.Group("/classes", mdw.IsValidToken)
	classes.GET("", classAPI.List)
	classes.GET("/:id", classAPI.Get)
	classes.POST("", classAPI.Create, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))
	classes.PUT("/:id", classAPI.Update, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))
	classes.DELETE("/:id", classAPI.Delete, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))
	classes.GET("/:id/students", classAPI.ListStudents)
	classes.GET("/:id/teachers", classAPI.ListTeachers)
	classes.PUT("/:class-id/students/:student-id", classAPI.AddStudent, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))
	classes.PUT("/:class-id/teachers/:teacher-id", classAPI.AddTeacher, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))
	classes.DELETE("/:class-id/students/:student-id", classAPI.RemoveStudent, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))
	classes.DELETE("/:class-id/teachers/:teacher-id", classAPI.RemoveTeacher, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))

	// Start listening
	server.Logger.Fatal(server.Start("127.0.0.1:8080"))
}

func initializeHTTPServer() *echo.Echo {
	// Echo instance customization
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Logger())  // Logger middleware
	e.Use(middleware.Recover()) // Panic recover middleware

	return e
}
