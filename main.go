package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// User Controller
func UserController(c echo.Context) error {
	id := c.Param("id")
	// Render Data - JSON Response
	return c.JSON(http.StatusOK, map[string]string{
		"id":   id,
		"name": "User name show here",
	})
}

func main() {
	e := echo.New()
	// routing
	e.GET("/users/:id", UserController)

	e.Start(":8000")
}
