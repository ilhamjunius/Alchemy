package routes

import (
	"github.com/alchemy/Day2/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	group := e.Group("v1")

	group.GET("/users", controllers.GetAllUsersController)
	group.GET("/users/:id", controllers.GetUserController)
	group.POST("/users", controllers.CreateUserController)
	group.DELETE("/users/:id", controllers.DeleteUserController)
	group.PUT("/users/:id", controllers.UpdateUserController)

	group.GET("/books", controllers.GetAllBooksController)
	group.GET("/books/:id", controllers.GetBookController)
	group.POST("/books", controllers.CreateBookController)
	group.DELETE("/books/:id", controllers.DeleteBookController)
	group.PUT("/books/:id", controllers.UpdateBookController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
	return e
}
