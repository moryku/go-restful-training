package routes

import (
	"github.com/labstack/echo"
	c "restful_training/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	// user routing
	e.GET("/api/users", c.GetUsersController)
	e.GET("/api/users/:id", c.GetUserController)
	e.POST("/api/users", c.CreateUserController)
	e.DELETE("/api/users/:id", c.DeleteUserController)
	e.PUT("/api/users/:id", c.UpdateUserController)

	// book routing
	e.GET("/api/books", c.GetBooksController)
	e.GET("/api/books/:id", c.GetBookController)
	e.GET("/api/books", c.GetBooksLikeController)
	e.POST("/api/books", c.CreateBookController)
	e.DELETE("/api/books/:id", c.DeleteBookController)
	e.PUT("/api/books/:id", c.UpdateBookController)

	return e
}
