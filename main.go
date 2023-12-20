package main

import (
	"net/http"
	"student-service/pkg/application/rest"
	dataaccess "student-service/pkg/data-access"
	"student-service/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	sqlDB := dataaccess.InitializeSequelDB("postgres://user:password@localhost:5432/student-service?sslmode=disable")

	// create student data access
	studentDA := dataaccess.NewStudentDA(sqlDB)

	// create student service
	studentService := service.NewStudentService(studentDA)

	// create student API
	studentAPI := rest.NewStudentAPI(studentService)

	server := initializeHTTPServer()

	// Index page
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	// Routes
	server.GET("/students", studentAPI.List)

	// server.GET("/login", example.Handle)

	// authenticated := server.Group("", authnMiddleware)
	// authenticated.GET("/me", example.Handle)

	// teacher
	// authenticated.GET("/classes", example.Handle)
	// authenticated.POST("/classes", example.Handle)
	// authenticated.PATCH("/classes/:classId", example.Handle)

	// admin
	// authenticated.GET("/users", example.Handle)
	// authenticated.POST("/users", example.Handle)
	// authenticated.PATCH("/users/:userId", example.Handle)

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
