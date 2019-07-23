package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func UserSearchController(c echo.Context) error {
	// get data from query param
	match := c.QueryParam("match")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"match":  match,
		"result": []string{"adi", "aan", "asif"}, // hardcode data
	})
}

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", UserSearchController)

	e.Start(":8000")
}
