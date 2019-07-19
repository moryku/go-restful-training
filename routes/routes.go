package routes

import (
	"github.com/iswanulumam/go-restful-training/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// routing
	e.GET("/api/hello", controllers.HelloController)
	e.GET("/api/users", controllers.FindAllUserController)
	e.POST("/api/users", controllers.CreateUserController)
	e.DELETE("/api/users/:id", controllers.DeleteUserController)

	return e
}
