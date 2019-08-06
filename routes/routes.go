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
	e.GET("/api/users/:id", c.GetUserController, m.BasicAuth(auth.BasicAuthCheckDB))
	e.GET("/api/users", c.GetUsersLikeControlller)
	e.POST("/api/users", c.CreateUserController)
	e.DELETE("/api/users/:id", c.DeleteUserController)
	e.PUT("/api/users/:id", c.UpdateUserController)

	return e
}
