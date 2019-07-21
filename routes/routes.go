package routes

import (
	"github.com/iswanulumam/go-restful-training/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// hello routing
	e.GET("/api/hello", controllers.HelloController)

	// user routing
	e.GET("/api/users", controllers.GetUsersController)
	e.GET("/api/users/:id", controllers.GetUserByIdController)
	e.GET("/api/users", controllers.GetUserLikeControlller)
	e.POST("/api/users", controllers.CreateUserController)
	e.DELETE("/api/users/:id", controllers.DeleteUserController)
	e.PUT("/api/users/:id", controllers.UpdateUserController)

	// book routing
	e.GET("/api/books", controllers.GetBooksController)
	e.GET("/api/books/:id", controllers.GetBookByIdController)
	e.GET("/api/books", controllers.GetBookLikeController)
	e.POST("/api/books", controllers.CreateBookController)
	e.DELETE("/api/books/:id", controllers.DeleteBookController)
	e.PUT("/api/books/:id", controllers.UpdateBookController)

	return e
}
