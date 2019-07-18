package routes

import (
	"github.com/iswanulumam/go-restful-training/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// routing
	e.GET("/api/hello", controllers.HelloController)

	return e
}
