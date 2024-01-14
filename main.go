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

	// create student data access
	studentDA := dataaccess.NewStudentDA(sqlDB)
	classDA := dataaccess.NewClassDA(sqlDB)
	userDA := dataaccess.NewUserDA(sqlDB)

	// create student service
	studentService := service.NewStudentService(studentDA)
	classService := service.NewClassService(classDA)
	userService := service.NewUserService(userDA)

	// create student API
	studentAPI := rest.NewStudentAPI(studentService)
	classAPI := rest.NewClassAPI(classService)
	userAPI := rest.NewUserAPI(userService)

	server := initializeHTTPServer()

	// Index page
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

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

	// Users
	users := server.Group("/users", mdw.IsValidToken)
	users.PUT("/:id", userAPI.Update, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin}))
	users.GET("/:id", userAPI.Get, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher, dto.Role_Student}))
	// server.GET("/login", example.Handle)
	server.POST("/login", handler.Login, mdw.BasicAuthWithUserService(userService))

	//Example using middleware for checking token and checking permission
	server.GET("/students", studentAPI.List, mdw.IsValidToken, mdw.IsValidPermission(userService, []dto.Role{dto.Role_Admin, dto.Role_Teacher}))

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
