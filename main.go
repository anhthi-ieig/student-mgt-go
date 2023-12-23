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
	classDA := dataaccess.NewClassDA(sqlDB)

	// create student service
	studentService := service.NewStudentService(studentDA)
	classService := service.NewClassService(classDA)

	// create student API
	studentAPI := rest.NewStudentAPI(studentService)
	classAPI := rest.NewClassAPI(classService)

	server := initializeHTTPServer()

	// Index page
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	// Routes
	server.GET("/students", studentAPI.List)

	// Classes
	server.GET("/classes", classAPI.List)
	server.GET("/classes/:id", classAPI.Get)
	server.POST("/classes", classAPI.Create)
	server.PUT("/classes/:id", classAPI.Update)
	server.DELETE("/classes/:id", classAPI.Delete)
	server.GET("/classes/:id/students", classAPI.ListStudents)
	server.GET("/classes/:id/teachers", classAPI.ListTeachers)
	server.PUT("/classes/:class-id/students/:student-id", classAPI.AddStudent)
	server.PUT("/classes/:class-id/teachers/:teacher-id", classAPI.AddTeacher)
	server.DELETE("/classes/:class-id/students/:student-id", classAPI.RemoveStudent)
	server.DELETE("/classes/:class-id/teachers/:teacher-id", classAPI.RemoveTeacher)

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
