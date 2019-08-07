package routes

import (
	c "github.com/iswanulumam/go-restful-training/controllers"
	auth "github.com/iswanulumam/go-restful-training/middlewares"

	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// user routing
	e.GET("/api/users", c.GetUsersController)
	e.GET("/api/users/:id", c.GetUserController)
	e.GET("/api/users", c.GetUsersLikeControlller)
	e.POST("/api/users", c.CreateUserController)
	e.DELETE("/api/users/:id", c.DeleteUserController, m.BasicAuth(auth.BasicAuth))
	e.PUT("/api/users/:id", c.UpdateUserController, m.BasicAuth(auth.BasicAuth))

	// book routing
	e.GET("/api/books", c.GetBooksController, m.BasicAuth(auth.BasicAuth2))
	e.GET("/api/books/:id", c.GetBookController, m.BasicAuth(auth.BasicAuth2))
	e.GET("/api/books", c.GetBooksLikeController, m.BasicAuth(auth.BasicAuth2))
	e.POST("/api/books", c.CreateBookController, m.BasicAuth(auth.BasicAuth2))
	e.DELETE("/api/books/:id", c.DeleteBookController, m.BasicAuth(auth.BasicAuth2))
	e.PUT("/api/books/:id", c.UpdateBookController, m.BasicAuth(auth.BasicAuth2))

	return e
}
