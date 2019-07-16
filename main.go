package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func helloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from the webiste")
}

func main() {
	// create a new echo instance
	e := echo.New()

	// implement middleware logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// routing
	e.GET("/", helloController)

	e.Start(":8081")
}
