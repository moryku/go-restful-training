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
	e.GET("/api/users", controllers.FindAllUserController)
	e.GET("/api/users/:id", controllers.FindUserByIdController)
	e.GET("/api/users", controllers.FindUserByLikeControlller)
	e.POST("/api/users", controllers.CreateUserController)
	e.DELETE("/api/users/:id", controllers.DeleteUserController)
	e.PUT("/api/users/:id", controllers.UpdateUserController)

	return e
}
