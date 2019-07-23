package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string
	Email string
}

func GetUser(c echo.Context) error {
	user := User{Name: "Ismail", Email: "ismail@alterra.id"}
	return c.JSON(http.StatusOK, user)
}

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/user", GetUser)
	// start the server, and log if it fails
	e.Start(":8000")
}
