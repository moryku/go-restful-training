package routes

import (
	c "github.com/iswanulumam/go-restful-training/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// hello routing
	e.GET("/api/hello", c.HelloController)

	// user routing
	e.GET("/api/users", c.GetUsersController)
	e.GET("/api/users/:id", c.GetUserByIdController)
	e.GET("/api/users", c.GetUserLikeControlller)
	e.POST("/api/users", c.CreateUserController)
	e.DELETE("/api/users/:id", c.DeleteUserController)
	e.PUT("/api/users/:id", c.UpdateUserController)

	// book routing
	e.GET("/api/books", c.GetBooksController)
	e.GET("/api/books/:id", c.GetBookByIdController)
	e.GET("/api/books", c.GetBookLikeController)
	e.POST("/api/books", c.CreateBookController)
	e.DELETE("/api/books/:id", c.DeleteBookController)
	e.PUT("/api/books/:id", c.UpdateBookController)

	return e
}
