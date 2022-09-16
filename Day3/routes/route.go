package routes

import (
	"github.com/alchemy/Day3/configs"
	"github.com/alchemy/Day3/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	group := e.Group("v1")
	auth := e.Group("v1")
	auth.Use(middleware.JWT([]byte(configs.JWT_SECRET)))

	auth.GET("/users", controllers.GetAllUsersController)
	auth.GET("/users/:id", controllers.GetUserController)
	group.POST("/users", controllers.CreateUserController)
	auth.DELETE("/users/:id", controllers.DeleteUserController)
	auth.PUT("/users/:id", controllers.UpdateUserController)
	group.POST("/login", controllers.Login)

	group.GET("/books", controllers.GetAllBooksController)
	group.GET("/books/:id", controllers.GetBookController)
	auth.POST("/books", controllers.CreateBookController)
	auth.DELETE("/books/:id", controllers.DeleteBookController)
	auth.PUT("/books/:id", controllers.UpdateBookController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
	return e
}
