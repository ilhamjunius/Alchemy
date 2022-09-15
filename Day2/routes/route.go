package routes

import (
	"Alterra/batch5/ORM/Part1/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	// Route / to handler function
	eAuth := e.Group("v1")
	// // eAuth.Use(mi)
	// eAuth.Use(middleware.BasicAuth(custMiddlewares.DBBasicAuth))
	// eAuth := e.Group("/akses")
	// eAuth.Use(middleware.JWT([]byte(configs.JWT_SECRET)))
	// eAuth := e.Group("/akses")
	// eAuth.Use(middleware.JWT([]byte(configs.JWT_SECRET)))
	eAuth.GET("/users", controllers.GetAllUsersController)
	eAuth.GET("/users/:id", controllers.GetUserController)
	eAuth.POST("/users", controllers.CreateUserController)
	eAuth.DELETE("/users/:id", controllers.DeleteUserController)
	eAuth.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetAllBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	eAuth.POST("/books", controllers.CreateBookController)
	eAuth.DELETE("/books/:id", controllers.DeleteBookController)
	eAuth.PUT("/books/:id", controllers.UpdateBookController)
	e.POST("/login", controllers.Login)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
	return e
}
