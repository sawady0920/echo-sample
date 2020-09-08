package router

import (
	"echo-sample/interface/controller"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	// Echo instance
	e := echo.New()

	userController := controller.NewUserController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/", "hello")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo World!!!!!!")
	})
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/user", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/user", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/user/:id", func(c echo.Context) error { return userController.Delete(c) })

	todoController := controller.NewTodoController(NewSqlHandler())
	e.GET("/todos", func(c echo.Context) error { return todoController.Index(c) })
	e.GET("/todo/:id", func(c echo.Context) error { return todoController.Show(c) })
	e.POST("todo", func(c echo.Context) error { return todoController.Create(c) })
	e.PUT("/todo", func(c echo.Context) error { return todoController.Save(c) })
	e.DELETE("/todo/:id", func(c echo.Context) error { return todoController.Delete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
