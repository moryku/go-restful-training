package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HelloController(c echo.Context) error {
	name := c.QueryParam("name")
	messages := "Hello " + name + "!"
	return c.JSON(http.StatusOK, map[string]string{
		"messages": messages,
	})
}
